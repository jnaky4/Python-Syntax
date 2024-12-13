package sse

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	sendMessage := func(data string) {
		fmt.Fprintf(w, "data: %s\n\n", data)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			sendMessage(fmt.Sprintf("Current time: %s", t.Format(time.RFC3339)))
		case <-r.Context().Done():
			return // Connection closed
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", sseHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  0, // Disable idle timeout
		ReadTimeout:  0, // Disable read timeout
		WriteTimeout: 0, // Disable write timeout
	}

	fmt.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}

//package main
//
//import (
//	"fmt"
//	"net/http"
//	"sync"
//	"time"
//)
//
//// LoggerState represents different states of the logger
//type LoggerState int
//
//const (
//	Muted LoggerState = iota
//	Logging
//	Triage
//)
//
//// Logger struct represents the logger with state management and log streaming
//type Logger struct {
//	state     LoggerState
//	mu        sync.Mutex
//	logStream []chan string // Channels to distribute log messages
//}
//
//// NewLogger creates a new Logger instance with initial state
//func NewLogger() *Logger {
//	return &Logger{
//		state:     Muted,
//		logStream: []chan string{},
//	}
//}
//
//// Start starts the logger if it's in the Muted state
//func (l *Logger) Start() {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	if l.state == Muted {
//		l.state = Logging
//		l.broadcastLog("Logger started: Logging all")
//	}
//}
//
//// Stop stops the logger and sets it to Muted
//func (l *Logger) Stop() {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	if l.state == Logging || l.state == Triage {
//		l.state = Muted
//		l.broadcastLog("Logger stopped")
//	}
//}
//
//// SetState sets the logger state and handles special cases for Triage
//func (l *Logger) SetState(state LoggerState) {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//
//	switch state {
//	case Triage:
//		if l.state == Muted {
//			l.state = Triage
//			l.broadcastLog("Logger set to Triage: Temporary logging activated")
//			go func() {
//				time.Sleep(10 * time.Second) // Adjust duration as needed
//				l.Stop()                     // Automatically stop logging after some time
//			}()
//		}
//	case Logging:
//		l.Start()
//	case Muted:
//		l.Stop()
//	}
//}
//
//// broadcastLog sends a log message to all logStream channels
//func (l *Logger) broadcastLog(message string) {
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	for _, ch := range l.logStream {
//		ch <- message
//	}
//}
//
//// HTTP handler for processing commands and state changes
//func (l *Logger) eventHandler(w http.ResponseWriter, r *http.Request) {
//	clientIP := r.RemoteAddr
//	fmt.Printf("Received event from IP: %s\n", clientIP)
//	fmt.Printf("Request Method: %s\n", r.Method)
//	fmt.Printf("Request URI: %s\n", r.RequestURI)
//
//	if r.Method != http.MethodPost {
//		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
//		return
//	}
//
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, "Failed to parse request", http.StatusBadRequest)
//		return
//	}
//
//	action := r.FormValue("action")
//	stateStr := r.FormValue("state")
//
//	switch action {
//	case "start":
//		l.Start()
//	case "stop":
//		l.Stop()
//	case "set":
//		var state LoggerState
//		switch stateStr {
//		case "logging":
//			state = Logging
//		case "triage":
//			state = Triage
//		default:
//			http.Error(w, "Invalid state", http.StatusBadRequest)
//			return
//		}
//		l.SetState(state)
//	default:
//		http.Error(w, "Invalid action", http.StatusBadRequest)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
//
//// HTTP handler for Server-Sent Events (SSE)
//func (l *Logger) sseHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/event-stream")
//	w.Header().Set("Cache-Control", "no-cache")
//	w.Header().Set("Connection", "keep-alive")
//
//	logChan := make(chan string)
//	l.mu.Lock()
//	l.logStream = append(l.logStream, logChan)
//	l.mu.Unlock()
//
//	clientIP := r.RemoteAddr
//	fmt.Printf("Client connected to SSE endpoint from IP: %s\n", clientIP)
//
//	defer func() {
//		l.mu.Lock()
//		for i, ch := range l.logStream {
//			if ch == logChan {
//				l.logStream = append(l.logStream[:i], l.logStream[i+1:]...)
//				break
//			}
//		}
//		l.mu.Unlock()
//		close(logChan)
//		fmt.Printf("Client disconnected from SSE endpoint from IP: %s\n", clientIP)
//	}()
//
//	for {
//		select {
//		case <-r.Context().Done():
//			return
//		case log, ok := <-logChan:
//			if !ok {
//				return
//			}
//			if _, err := fmt.Fprintf(w, "data: %s\n\n", log); err != nil {
//				return
//			}
//			if f, ok := w.(http.Flusher); ok {
//				f.Flush()
//			}
//		}
//	}
//}
//
//func main() {
//	logger := NewLogger()
//
//	http.HandleFunc("/event", logger.eventHandler)
//	http.HandleFunc("/sse", logger.sseHandler)
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		fmt.Println("Server failed:", err)
//	}
//}
//
////package main
////
////import (
////	"encoding/json"
////	"fmt"
////	"net/http"
////	"sync"
////	"time"
////)
////
////// LoggerState represents different states of the logger
////type LoggerState int
////
////const (
////	Muted LoggerState = iota
////	Logging
////	Triage
////)
////
////// EventType represents different types of events
////type EventType int
////
////const (
////	StartLogging EventType = iota
////	StopLogging
////	ChangeState
////)
////
////// Event struct represents an event with type and optional data
////type Event struct {
////	Type EventType
////	Data interface{}
////}
////
////// Logger struct represents the logger with state management and log streaming
////type Logger struct {
////	state     LoggerState
////	mu        sync.Mutex
////	logStream []chan string // Channels to distribute log messages
////}
////
////// NewLogger creates a new Logger instance with initial state
////func NewLogger() *Logger {
////	return &Logger{
////		state:     Muted,
////		logStream: []chan string{},
////	}
////}
////
////// Start starts the logger if it's in the Muted state
////func (l *Logger) Start() {
////	l.mu.Lock()
////	defer l.mu.Unlock()
////	if l.state == Muted {
////		l.state = Logging
////		message := "Logger started: Logging all"
////		fmt.Println(message)
////		l.broadcastLog(message)
////	}
////}
////
////// Stop stops the logger and sets it to Muted
////func (l *Logger) Stop() {
////	l.mu.Lock()
////	defer l.mu.Unlock()
////	if l.state == Logging || l.state == Triage {
////		l.state = Muted
////		message := "Logger stopped"
////		fmt.Println(message)
////		l.broadcastLog(message)
////	}
////}
////
////// SetState sets the logger state and handles special cases for Triage
////func (l *Logger) SetState(state LoggerState) {
////	l.mu.Lock()
////	defer l.mu.Unlock()
////
////	switch state {
////	case Triage:
////		if l.state == Muted {
////			l.state = Triage
////			message := "Logger set to Triage: Temporary logging activated"
////			fmt.Println(message)
////			l.broadcastLog(message)
////			go func() {
////				time.Sleep(10 * time.Second) // Adjust duration as needed
////				l.Stop()                     // Automatically stop logging after some time
////			}()
////		}
////	case Logging:
////		l.Start()
////	case Muted:
////		l.Stop()
////	}
////}
////
////// AutoRunner listens for events and handles logger state changes
////func (l *Logger) AutoRunner(eventCh <-chan Event, wg *sync.WaitGroup) {
////	defer wg.Done()
////	for event := range eventCh {
////		switch event.Type {
////		case StartLogging:
////			l.Start()
////		case StopLogging:
////			l.Stop()
////		case ChangeState:
////			if state, ok := event.Data.(LoggerState); ok {
////				l.SetState(state)
////			}
////		}
////	}
////}
////
////// broadcastLog sends a log message to all logStream channels
////func (l *Logger) broadcastLog(message string) {
////	l.mu.Lock()
////	defer l.mu.Unlock()
////	for _, ch := range l.logStream {
////		ch <- message
////	}
////}
////
////// HTTP handler for processing commands
////func (l *Logger) eventHandler(w http.ResponseWriter, r *http.Request) {
////	clientIP := r.RemoteAddr
////	fmt.Printf("Received event from IP: %s\n", clientIP)
////	fmt.Printf("Request Method: %s\n", r.Method)
////	fmt.Printf("Request URI: %s\n", r.RequestURI)
////
////	var event Event
////	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
////		http.Error(w, "Invalid request payload", http.StatusBadRequest)
////		return
////	}
////
////	switch event.Type {
////	case StartLogging:
////		l.Start()
////	case StopLogging:
////		l.Stop()
////	case ChangeState:
////		if state, ok := event.Data.(LoggerState); ok {
////			l.SetState(state)
////		}
////	}
////
////	w.WriteHeader(http.StatusOK)
////}
////
////// HTTP handler for Server-Sent Events (SSE)
////func (l *Logger) sseHandler(w http.ResponseWriter, r *http.Request) {
////	w.Header().Set("Content-Type", "text/event-stream")
////	w.Header().Set("Cache-Control", "no-cache")
////	w.Header().Set("Connection", "keep-alive")
////
////	logChan := make(chan string)
////	l.mu.Lock()
////	l.logStream = append(l.logStream, logChan)
////	l.mu.Unlock()
////
////	clientIP := r.RemoteAddr
////	fmt.Printf("Client connected to SSE endpoint from IP: %s\n", clientIP)
////
////	defer func() {
////		l.mu.Lock()
////		for i, ch := range l.logStream {
////			if ch == logChan {
////				l.logStream = append(l.logStream[:i], l.logStream[i+1:]...)
////				break
////			}
////		}
////		l.mu.Unlock()
////		close(logChan)
////		fmt.Printf("Client disconnected from SSE endpoint from IP: %s\n", clientIP)
////	}()
////
////	for {
////		select {
////		case <-r.Context().Done():
////			fmt.Println("Client connection closed")
////			return
////		case log, ok := <-logChan:
////			if !ok {
////				fmt.Println("Log channel closed")
////				return
////			}
////			if _, err := fmt.Fprintf(w, "data: %s\n\n", log); err != nil {
////				fmt.Println("Error writing to SSE client:", err)
////				return
////			}
////			if f, ok := w.(http.Flusher); ok {
////				f.Flush()
////			}
////		}
////	}
////}
////
////func main() {
////	logger := NewLogger()
////	eventCh := make(chan Event)
////	var wg sync.WaitGroup
////	wg.Add(1)
////	go logger.AutoRunner(eventCh, &wg)
////
////	http.HandleFunc("/event", logger.eventHandler)
////	http.HandleFunc("/sse", logger.sseHandler)
////	if err := http.ListenAndServe(":8080", nil); err != nil {
////		fmt.Println("Server failed:", err)
////	}
////
////	close(eventCh)
////	wg.Wait()
////}
