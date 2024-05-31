package main

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net"
	"sync"
	"syscall"
	"time"
)

const (
	SERVER_HOST   = "localhost"
	SERVER_PORT   = "9988"
	SERVER_TYPE   = "tcp"
	SERVER_SOCKET = SERVER_HOST + ":" + SERVER_PORT
	PING_INTERVAL = 30 * time.Second
)

func main() {
	var wg sync.WaitGroup
	messages := [...]string{"Jake", "Sam", "Andrew", "Taylor"}
	for i := range messages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			SendAMessage(messages[i])
		}()

	}
	wg.Wait()
}

func SendAMessage(message string) {
	connection, err := net.Dial(SERVER_TYPE, SERVER_SOCKET)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v\n", connection.LocalAddr())

	_, err = connection.Write([]byte(message))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}

// DialTimeout Wraps a Connection Interface to set the timeout
func DialTimeout(network, address string, timeout time.Duration,
) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

// DialContextTimeout Uses Context to set timeout
func DialContextTimeout() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	var d net.Dialer
	socket, err := d.DialContext(ctx, SERVER_TYPE, SERVER_SOCKET)
	if err != nil {
		socket.Close()
		println(err)
	}

	//capturing the Timeout
	if ctx.Err() == context.DeadlineExceeded {
		println("Deadline Exceeded")
	}

	//Message logic here
}

func DialWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	//sync := make(chan struct{})
	var d net.Dialer
	socket, err := d.DialContext(ctx, SERVER_TYPE, SERVER_SOCKET)
	if err != nil {
		return
	}
	fmt.Printf("%v/n", socket.LocalAddr())
	cancel()

}

//Canceling Multiple Sockets
//You can pass the same context to multiple DialContext
//calls and cancel all the calls at the same time by
//executing the context’s cancel function. For example,
//let’s assume you need to retrieve a resource via TCP
//that is on several servers. You can asynchronously
//dial each server, passing each dialer the same context.
//You can then abort the remaining dialers after you
//receive a response from one of the servers.

func DialMultipleWithCancel() {
	//todo
}

// implementing a heartbeat
func HeartBeat(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset: //1 pulled initial interval off reset channel
	default:
	}
	if interval <= 0 {
		interval = PING_INTERVAL
	}

	timer := time.NewTimer(interval) //2
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done(): //3
			return
		case newInterval := <-reset: //4
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C: //5
			if _, err := w.Write([]byte("ping")); err != nil {
				// track and act on consecutive timeouts here
				return
			}
		}

		_ = timer.Reset(interval)
	}
}

//The Heartbeat function writes ping messages to a given writer at
//regular intervals. Because it’s meant to run in a goroutine,
//Pinger accepts a context as its first argument so you can
//terminate it and prevent it from leaking. Its remaining arguments
//include an io.Writer interface and a channel to signal a timer reset.
//You create a buffered channel and put a duration on it to set the timer’s
//initial interval 1. If the interval isn’t greater than zero, you use the
//default ping interval.
//
//You initialize the timer to the interval 2 and set up a deferred
//call to drain the timer’s channel to avoid leaking it, if necessary.
//The endless for loop contains a select statement, where you block
//until one of three things happens: the context is canceled, a
//signal to reset the timer is received, or the timer expires.
//If the context is canceled 3, the function returns, and no
//further pings will be sent. If the code selects the reset channel
//4, you shouldn’t send a ping, and the timer resets 6 before
//iterating on the select statement again.
//
//If the timer expires 5, you write a ping message to the writer,
//and the timer resets before the next iteration. If you wanted,
//you could use this case to keep track of any consecutive time-outs
//that occur while writing to the writer. To do this, you could pass
//in the context’s cancel function and call it here if you reach a
//threshold of consecutive time-outs.

func ExamplePinger() {
	ctx, cancel := context.WithCancel(context.Background())
	r, w := io.Pipe() // in lieu of net.Conn
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second // initial ping interval

	go func() {
		HeartBeat(ctx, w, resetTimer)
		close(done)
	}()

	receivePing := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			fmt.Printf("resetting timer (%s)\n", d)
			resetTimer <- d
		}

		now := time.Now()
		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("received %q (%s)\n",
			buf[:n], time.Since(now).Round(100*time.Millisecond))
	}

	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d:\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}

	cancel()
	<-done // ensures the pinger exits after canceling the context

	// Output:
	// Run 1:
	// resetting timer (0s)
	// received "ping" (1s)
	// Run 2:
	// resetting timer (200ms)
	// received "ping" (200ms)
	// Run 3:
	// resetting timer (300ms)
	// received "ping" (300ms)
	// Run 4:
	// resetting timer (0s)
	// received "ping" (300ms)
	// Run 5:
	// received "ping" (300ms)
	// Run 6:
	// received "ping" (300ms)
	// Run 7:
	// received "ping" (300ms)

}

/*
In this example, you create a buffered channel 1 that you’ll use
to signal a reset of the Pinger’s timer. You put an initial ping
interval of one second on the resetTimer channel before passing
the channel to the Pinger function. You’ll use this duration to
initialize the Pinger’s timer and dictate when to write the ping
message to the writer.

You run through a series of millisecond durations in a loop 2,
passing each to the receivePing function. This function resets
the ping timer to the given duration and then waits to receive
the ping message on the given reader. Finally, it prints to stdout
the time it takes to receive the ping message. Go checks stdout
against the expected output in the example.

During the first iteration 3, you pass in a duration of zero,
which tells the Pinger to reset its timer by using the previous
duration—one second in this example. As expected, the reader
receives the ping message after one second. The second iteration
4 resets the ping timer to 200 ms. Once this expires, the reader
receives the ping message. The third run resets the ping timer to
300 ms 5, and the ping arrives at the 300 ms mark.

You pass in a zero duration for run 4 6, preserving the 300 ms
ping timer from the previous run. I find the technique of using
zero durations to mean “use the previous timer duration” useful
because I do not need to keep track of the initial ping timer
duration. I can simply initialize the timer with the duration
I want to use for the remainder of the TCP session and reset
the timer by passing in a zero duration every time I need to
preempt the transmission of the next ping message. Changing
the ping timer duration in the future involves the modification
of a single line as opposed to every place I send on the resetTimer
channel.

Runs 5 to 7 7 simply listen for incoming pings without resetting
the ping timer. As expected, the reader receives a ping at 300
ms intervals for the last three runs
*/

/* advancing the Deadline by Using the Heartbeat
Each side of a network connection could use a Pinger to advance
its deadline if the other side becomes idle, whereas the previous
examples showed only a single side using a Pinger. When either node
receives data on the network connection, its ping timer should reset
to stop the delivery of an unnecessary ping.
*/
