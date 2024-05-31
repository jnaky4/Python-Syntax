package main

import (
	"context"
	"fmt"
"net/http"
"time"
)
//req.Context() vs context.Background


func ctxPrint(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: ctxPrint handler started")
	defer fmt.Println("server: ctxPrint handler ended")


	select {
		//Wait for a few seconds before sending a reply to the client. This could simulate some work the server is doing.
		//While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and
		//return as soon as possible.
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "context %+v\n", ctx)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}


func ctxTimeout(w http.ResponseWriter, req *http.Request){
	ctx := req.Context()

	deadline := time.Now().Add(1 * time.Second)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()
}

func main() {
	//context.Background returns a non-nil, empty [Context]. It is never canceled, has no values, and has no deadline.
	//It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming
	//requests
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	//context.WithCancel returns a copy of parent with a new Done channel. The returned context's Done channel is closed when
	//the returned cancel function is called or when the parent context's Done channel is closed, whichever happens
	//first. Canceling this context releases resources associated with it, so code should call cancel as soon as the
	//operations running in this Context complete

	http.HandleFunc("/ctx", ctxPrint)
	http.HandleFunc("/ctx/timeout", ctxTimeout)

	http.ListenAndServe(":8090", nil)
}

