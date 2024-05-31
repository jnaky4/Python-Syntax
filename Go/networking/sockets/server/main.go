package main

import (
	"fmt"
	"net"
	"os"
	"testing"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func TestListener(t *testing.T) {
	//net.Listen accepts a network type: tcp or udp
	//socket: ip:port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	/*
		Omitting the port or setting to 0 will randomly assign a port number to the listener
		Omitting the IP address, the listener is bound to all unicast and anycast IP addresses
		You can restrict the listener to IPv4 addressed by passing tcp4 and IPv6 with tcp6

	*/
	if err != nil { //if no error, the listener is bound to the socket
		t.Fatal(err)
	}
	//always gracefully shut down
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())

}

func main() {
	fmt.Println("Server Starting...")
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)

	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}
