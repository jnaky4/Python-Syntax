package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	TaskID int
	Output interface{}
}

type WorkerPool struct {
	numWorkers int
	tasks      chan Task
	results    chan Result
	wg         sync.WaitGroup
}

type ResultWorkerPool struct {
	numWorkers int
	results    chan Result
	wg         sync.WaitGroup
}

type Task struct {
	ID   int
	Func interface{}
	Arg  interface{}
}

type GeneralWorkerPool struct {
	numWorkers int
	tasks      chan Task
	results    chan Result
	wg         sync.WaitGroup
}

func NewGeneralWorkerPool(numWorkers int) *GeneralWorkerPool {
	return &GeneralWorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan Task, 100),
		results:    make(chan Result, 100),
	}
}

func (wp *GeneralWorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		go wp.worker(i)
	}
}

func (wp *GeneralWorkerPool) worker(id int) {
	for task := range wp.tasks {
		var output interface{}
		// Handle the task based on its function type dynamically
		switch fn := task.Func.(type) {
		case func(int) string:
			output = fn(task.Arg.(int)) // Handle func(int) string
		case func(string) int:
			output = fn(task.Arg.(string)) // Handle func(string) int
		default:
			output = "Unknown function type"
		}
		wp.results <- Result{TaskID: task.ID, Output: output}
		wp.wg.Done()
	}
}

func (wp *GeneralWorkerPool) AddTask(task Task) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *GeneralWorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.tasks)
	close(wp.results)
}

func (rp *ResultWorkerPool) Start() {
	for i := 0; i < rp.numWorkers; i++ {
		go rp.resultWorker(i)
	}
}

func (rp *ResultWorkerPool) resultWorker(id int) {
	for result := range rp.results {
		// Process result
		fmt.Printf("Result from Task %d: %v\n", result.TaskID, result.Output)
		rp.wg.Done()
	}
}

func (rp *ResultWorkerPool) AddResult(result Result) {
	rp.wg.Add(1)
	rp.results <- result
}

func (rp *ResultWorkerPool) Wait() {
	rp.wg.Wait()
	close(rp.results)
}

// Function 1: Takes an int and returns a string
func exampleFunction1(n int) string {
	time.Sleep(1 * time.Second) // Simulate some work
	return fmt.Sprintf("Processed number: %d", n)
}

// Function 2: Takes a string and returns an int (length of the string)
func exampleFunction2(s string) int {
	time.Sleep(1 * time.Second) // Simulate some work
	return len(s)
}

func main() {
	// Create a worker pool for tasks that will work with any function type
	wp := NewGeneralWorkerPool(3) // 3 worker goroutines for tasks
	wp.Start()

	// Create a worker pool for results processing
	rp := &ResultWorkerPool{numWorkers: 2, results: make(chan Result, 100)}
	rp.Start()

	// Add tasks to the worker pool (each with different function types)
	wp.AddTask(Task{ID: 1, Func: exampleFunction1, Arg: 10})      // Function 1: int -> string
	wp.AddTask(Task{ID: 2, Func: exampleFunction2, Arg: "test"})  // Function 2: string -> int
	wp.AddTask(Task{ID: 3, Func: exampleFunction1, Arg: 25})      // Function 1: int -> string
	wp.AddTask(Task{ID: 4, Func: exampleFunction2, Arg: "hello"}) // Function 2: string -> int

	// Wait for tasks to finish
	wp.Wait()

	// Now that tasks are done, process the results concurrently
	rp.Wait()

	fmt.Println("All tasks and results have been processed.")
}
