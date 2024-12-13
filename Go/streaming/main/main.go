package main

import (
	"Go/streaming"
	"Go/time_completion"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	tracker := time_completion.NewTimerTracker()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	tracker.Track("stream", func() { http.HandleFunc("/stream-book", streaming.StreamBookHandler) })
	http.ListenAndServe(":8080", nil)

	<-signalChannel

	defer tracker.Report()
}
