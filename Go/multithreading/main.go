package main

import (
	"Go/cli/colors"
	tc "Go/const/terminalColors"
	"Go/time_completion"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"syscall"
)

//type WorkerPool struct {
//	numWorkers int
//	tasks      chan func() // channel to hold tasks (functions)
//	wg         sync.WaitGroup
//}
//
//func NewWorkerPool(numWorkers int, channelSize int) *WorkerPool {
//	return &WorkerPool{
//		numWorkers: numWorkers,
//		tasks:      make(chan func(), channelSize),
//	}
//}
//
//func (wp *WorkerPool) Start() {
//	// Create a worker pool
//	for i := 0; i < wp.numWorkers; i++ {
//		go wp.worker(i)
//	}
//}
//
//func (wp *WorkerPool) worker(id int) {
//	for task := range wp.tasks {
//		//fmt.Printf("Worker %d started\n", id)
//		task() // Execute the task
//		//fmt.Printf("Worker %d finished\n", id)
//		wp.wg.Done()
//	}
//}
//
//func (wp *WorkerPool) AddTask(task func()) {
//	wp.wg.Add(1)
//	wp.tasks <- task // Send task to worker
//}
//
//func (wp *WorkerPool) Wait() {
//	wp.wg.Wait() // Wait for all tasks to be completed
//}

/*
Can the Worker Pool Handle Any Function Type?
In Go, all functions passed to the worker pool must share the same signature (in this case, func()—a function that returns nothing and takes no arguments). If you want to pass functions with different signatures (for example, functions that take parameters), you can use interfaces or function signatures as arguments.

For example, you could define a worker pool that accepts tasks of different types by using a Task interface and having all tasks implement that interface.

Conclusion:
This approach allows you to use a generic worker pool where the tasks can be any function, and the workers can process them concurrently.

*/

//func main() {
//	// Create a worker pool with 3 workers
//	wp := NewWorkerPool(3)
//
//	// Start the worker pool
//	wp.Start()
//
//	// Define some tasks (functions)
//	task1 := func() {
//		fmt.Println("Executing task 1")
//	}
//	task2 := func() {
//		fmt.Println("Executing task 2")
//	}
//	task3 := func() {
//		fmt.Println("Executing task 3")
//	}
//	task4 := func() {
//		fmt.Println("Executing task 4")
//	}

//	// Add tasks to the pool
//	wp.AddTask(task1)
//	wp.AddTask(task2)
//	wp.AddTask(task3)
//	wp.AddTask(task4)
//
//	// Wait for all tasks to be processed
//	wp.Wait()
//
//	fmt.Println("All tasks completed.")
//}

// todo implement generic worker pools to the functions
// todo continue optimizing

func main() {
	//oldMain()
	//newMain()
	//newMainWorkPool()
	testMain()
}

func testMain() {
	resultCh := make(chan Result, 10) // Buffered channel to avoid blocking

	filePath := "/Users/Z004X7X/Git/syntax/Go/multithreading/books/theLordOfTheRingsTrilogy.txt"
	searchStr := "time"

	// Start the search
	go func() {
		mmapBMSearchFile(filePath, searchStr, resultCh)
		close(resultCh)
	}()

	// Collect and print results
	for result := range resultCh {
		fmt.Printf("Path: %s\n", result.path)
		for loc, content := range result.location {
			fmt.Printf("Chunk: %d, Content: %s\n", loc, content)
		}
	}

	fmt.Println("Search complete.")
}

//type WorkerPool struct {
//	numWorkers int
//	wg         sync.WaitGroup
//}
//
//func NewWorkerPool(numWorkers int) *WorkerPool {
//	return &WorkerPool{
//		numWorkers: numWorkers,
//	}
//}
//
//func (wp *WorkerPool) AddTask(task func()) {
//	wp.wg.Add(1)
//	go func() {
//		defer wp.wg.Done()
//		task()
//	}()
//}
//
//func (wp *WorkerPool) Wait() {
//	wp.wg.Wait()
//}

type WorkerPool struct {
	numWorkers int
	tasks      chan func()
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan func()),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		go func() {
			for task := range wp.tasks {
				task()
				wp.wg.Done()
			}
		}()
	}
}

func (wp *WorkerPool) AddTask(task func()) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) Wait() {
	close(wp.tasks) // Ensure all workers terminate gracefully
	wp.wg.Wait()
}

var pool = sync.Pool{
	New: func() interface{} {
		return make(map[int]string)
	},
}

type Result struct {
	path     string
	location map[int]string
}

var books []string
var verbose bool
var printMutex sync.Mutex

func newMainWorkPool() {
	defer time_completion.Timer()()
	defer fmt.Println("Total Time: ")

	//paths := []string{
	//	"/Users/Z004X7X/Git/syntax/Go/multithreading/books",
	//}
	//if err := loadBooks(paths); err != nil {
	//	fmt.Println("Error loading books:", err)
	//	return
	//}

	books = []string{"/Users/Z004X7X/Git/syntax/Go/multithreading/books/theLordOfTheRingsTrilogy.txt"}

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	//maxGoroutines := numCPU * 2

	workerPool := NewWorkerPool(1)
	workerPool.Start()

	searchStrings := []string{
		"i", "no", "the", "time", "unfortunately", "tom bombadil", "the lord of the rings", "aaaaa", "alalal", "everybody wants to rule the world",
	}

	//searchStrings := []string{"time"}

	repeat := 5
	verbose = false

	tracker := time_completion.NewTimerTracker()

	resultCh := make(chan Result, len(books)*len(searchStrings)*repeat)
	defer close(resultCh)

	for i := 0; i < repeat; i++ {
		for _, searchStr := range searchStrings {
			for _, book := range books {

				searchStr := searchStr // Ensure closure captures the current value
				book := book

				workerPool.AddTask(func() {
					tracker.Track(
						fmt.Sprintf("%s -> %s", colors.SetColor(searchStr, tc.Salmon), colors.SetColor(filepath.Base(book), tc.Mint)),
						func() {
							mmapBMSearchFile(book, searchStr, resultCh)
						},
					)
				})

			}
		}
	}

	workerPool.Wait()
	//go GatherResults(resultCh)
	tracker.Report()
	//GatherResults(resultCh)
}

func mmapBMSearchFile(path string, searchStr string, resultCh chan<- Result) {
	// Open file for reading
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
		return
	}

	defer file.Close()

	// Get file size and memory-map the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %s, error: %v\n", path, err)
		return
	}
	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileInfo.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("Failed to memory-map file: %s, error: %v\n", path, err)
		return
	}

	defer syscall.Munmap(data)

	// Preprocess for Boyer-Moore
	searchBytes := []byte(strings.ToLower(searchStr))
	badChar := preprocessBadCharBM(searchBytes)

	// Use chunked memory processing rather than line-by-line
	chunkSize := 8 * 1024 // 8 KB
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		if boyerMooreMatchBM(chunk, searchBytes, badChar) {
			location := pool.Get().(map[int]string) // Get reusable map from pool
			defer pool.Put(location)
			location[i/chunkSize] = string(chunk)
			select {
			case resultCh <- Result{path: path, location: location}: // Ensure non-blocking result send
			default:
				// Handle the case where the result channel is full (avoid blocking)
				//log.Println("Result channel is full, skipping this result")
			}
		}
	}
}

// todo some questions, why are we chunking line by line? would a pointer to a large str be faster?
//func mmapBMSearchFile(path string, searchStr string, resultCh chan<- Result) {
//
//	// Open file for reading
//	file, err := os.Open(path)
//	if err != nil {
//		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
//		return
//	}
//	defer file.Close()
//
//	// Get file size and memory-map the file
//	fileInfo, err := file.Stat()
//	if err != nil {
//		fmt.Printf("Failed to get file info: %s, error: %v\n", path, err)
//		return
//	}
//	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileInfo.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
//	if err != nil {
//		fmt.Printf("Failed to memory-map file: %s, error: %v\n", path, err)
//		return
//	}
//
//	defer syscall.Munmap(data)
//
//	// Preprocess for Boyer-Moore
//	searchBytes := []byte(strings.ToLower(searchStr))
//	badChar := preprocessBadCharBM(searchBytes)
//	location := pool.Get().(map[int]string) // Get reusable map from pool
//	defer pool.Put(location)
//
//	// Search across the mapped data by line
//	lineStart, lineNum := 0, 1
//	for i := 0; i < len(data); i++ {
//		if data[i] == '\n' || i == len(data)-1 {
//			line := data[lineStart : i+1]
//			if boyerMooreMatchBM(line, searchBytes, badChar) {
//				location[lineNum] = string(line)
//			}
//			lineStart = i + 1
//			lineNum++
//		}
//	}
//
//	// Send results if any matches found
//	if len(location) > 0 {
//		resultCh <- Result{path: path, location: location}
//	}
//}

func oldMain() {
	defer time_completion.Timer()()
	defer print("Total Time: ")

	//All multithreading cases are required to be case in-sensitive, and search all files for matches of the string. None handle non-ascii chars

	paths := []string{
		"/Users/Z004X7X/Git/syntax/Go/multithreading/books",
	}

	if err := loadBooks(paths); err != nil {
		fmt.Println("Error loading books:", err)
		return
	}

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	maxGoroutines := numCPU * 2 // Limit number of concurrent goroutines with semaphore channel

	defer fmt.Println("Current GOMAXPROCS:", colors.SetColor(fmt.Sprintf("%d", runtime.GOMAXPROCS(0)), tc.Salmon))
	defer fmt.Printf("Number of CPUs available: %s\n", colors.SetColor(fmt.Sprintf("%d", numCPU), tc.Salmon))
	defer println("Book Count: ", colors.SetColor(fmt.Sprintf("%d", len(books)), tc.Salmon))
	defer println("Max Routines: ", colors.SetColor(fmt.Sprintf("%d", maxGoroutines), tc.Salmon))

	searchStrings := []string{
		"i", "no", "the", "time", "unfortunately", "tom bombadil", "the lord of the rings", "aaaaa", "alalal", "everybody wants to rule the world",
	}

	repeat := 10
	verbose = false

	for _, searchStr := range searchStrings {
		println("SEARCH STRING => ", colors.SetColor(searchStr, tc.Mint))

		//println(colors.SetColor("Single Thread", tc.ElectricBlue))
		//time_completion.TrackFunction(func() { SingleThreadedExample(searchStr) }, repeat)
		//println()

		//println(colors.SetColor("Multi Thread(1)", tc.ElectricBlue))
		//time_completion.TrackFunction(func() { basicMultithreadingExample(searchStr, 1) }, repeat)
		//println()

		//if verbose {
		//	println(colors.SetColor("Verbose", tc.ElectricBlue))
		//	basicMultithreadingExampleDisplayResults(searchStr, maxGoroutines)
		//	println()
		//}
		//
		//println(colors.SetColor("Strings.Contains", tc.ElectricBlue))
		//time_completion.TrackFunction(func() { basicMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		//println()

		//println("Marginal performance beyond 2 * ", numCPU, " = ", maxGoroutines, " count")
		//println("Max Routines: ", 100)
		//time_completion.TrackFunction(func() { basicMultithreadingExample(searchStr, 100) }, repeat)
		//println()

		//println("Max Routines: ", maxGoroutines)
		//time_completion.TrackFunction(func() { regexMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		//println()

		//time_completion.TrackFunction(func() { boyerMooreMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		//println()
		//
		//time_completion.TrackFunction(func() { mmapMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		//println()
		//
		//time_completion.TrackFunction(func() { mmapBMMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		//println()

		print("BEST CHOICE: ")
		switch { // all 3 implement some variant of boyer-moore
		case len(searchStr) < 2:
			println(colors.SetColor("Strings.Contains", tc.ElectricBlue))
			time_completion.TrackFunction(func() { basicMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		case len(searchStr) >= 10:
			println(colors.SetColor("MMAP Boyer-Moore", tc.ElectricBlue))
			time_completion.TrackFunction(func() { mmapBMMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		default:
			println(colors.SetColor("MMAP Strings.Contains", tc.ElectricBlue))
			time_completion.TrackFunction(func() { mmapMultithreadingExample(searchStr, maxGoroutines) }, repeat)
		}
		println()
	}
}

func newMain() {
	defer time_completion.Timer()()
	defer print("Total Time: ")

	paths := []string{
		"/Users/Z004X7X/Git/syntax/Go/multithreading/books",
	}

	if err := loadBooks(paths); err != nil {
		fmt.Println("Error loading books:", err)
		return
	}

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	maxGoroutines := numCPU * 2

	defer fmt.Println("Current GOMAXPROCS:", colors.SetColor(fmt.Sprintf("%d", runtime.GOMAXPROCS(0)), tc.Salmon))
	defer fmt.Printf("Number of CPUs available: %s\n", colors.SetColor(fmt.Sprintf("%d", numCPU), tc.Salmon))
	defer println("Book Count: ", colors.SetColor(fmt.Sprintf("%d", len(books)), tc.Salmon))
	defer println("Max Routines: ", colors.SetColor(fmt.Sprintf("%d", maxGoroutines), tc.Salmon))

	searchStrings := []string{
		"i", "no", "the", "time", "unfortunately", "tom bombadil", "the lord of the rings", "aaaaa", "alalal", "everybody wants to rule the world",
	}

	repeat := 25
	verbose = false

	tracker := time_completion.NewTimerTracker()

	//todo We set max go routines here, however the multithreading examples call their own threads.
	//todo create single thread instance of function
	workerPool := NewWorkerPool(maxGoroutines)

	for _, searchStr := range searchStrings {
		//println("SEARCH STRING => ", colors.SetColor(searchStr, tc.Mint))
		for i := 0; i < repeat; i++ {

			//print("BEST CHOICE: ")
			switch {
			case len(searchStr) < 2:
				workerPool.AddTask(func() {
					//println(colors.SetColor("Strings.Contains", tc.ElectricBlue))
					tracker.Track(
						fmt.Sprintf("basicMultithreadingExample-%s",
							colors.SetColor(searchStr, tc.Salmon)), func() {
							mmapBMMultithreadingExample(searchStr, maxGoroutines)
						})
				})
			case len(searchStr) >= 10:
				workerPool.AddTask(func() {
					//println(colors.SetColor("MMAP Boyer-Moore", tc.ElectricBlue))
					tracker.Track(
						fmt.Sprintf("mmapBMMultithreadingExample-%s", colors.SetColor(searchStr, tc.Salmon)),
						func() {
							mmapBMMultithreadingExample(searchStr, maxGoroutines)
						})
				})
			default:
				workerPool.AddTask(func() {
					//println(colors.SetColor("MMAP Strings.Contains", tc.ElectricBlue))
					tracker.Track(
						fmt.Sprintf("mmapMultithreadingExample-%s", colors.SetColor(searchStr, tc.Salmon)),
						func() {
							mmapBMMultithreadingExample(searchStr, maxGoroutines)
						})
				})
			}
		}
		workerPool.Wait()
	}
	tracker.Report()
}

func loadBooks(paths []string) error {
	var files []string
	for _, path := range paths {
		err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, p)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("error reading path %s: %w", path, err)
		}
	}
	books = files
	return nil
}

func basicSearchFile(path string, searchStr string, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
		return
	}
	defer file.Close()

	location := make(map[int]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0
	searchStr = strings.ToLower(searchStr) // Lowercase once if case-insensitive matching is needed

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Avoid repeated ToLower calls and memory allocations
		if strings.Contains(line, searchStr) || strings.EqualFold(line, searchStr) {
			location[lineNum] = line
		}
	}

	if len(location) > 0 {
		resultCh <- Result{path: path, location: location}
	}
}

func mmapSearchFile(path string, searchStr string, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
		return
	}
	defer file.Close()

	// Memory-map the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %s, error: %v\n", path, err)
		return
	}
	fileSize := fileInfo.Size()
	data, err := syscall.Mmap(int(file.Fd()), 0, int(fileSize), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("Failed to memory-map file: %s, error: %v\n", path, err)
		return
	}
	defer syscall.Munmap(data)

	// Lowercase search string once
	searchStrLower := []byte(strings.ToLower(searchStr))
	searchLen := len(searchStrLower)
	location := make(map[int]string)
	lineStart := 0
	lineNum := 1

	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || i == len(data)-1 {
			line := data[lineStart : i+1]

			// Use byte comparison for case-insensitive search within mmap data
			if containsIgnoreCase(line, searchStrLower, searchLen) {
				location[lineNum] = string(line)
			}

			lineStart = i + 1
			lineNum++
		}
	}

	if len(location) > 0 {
		resultCh <- Result{path: path, location: location}
	}
}

func containsIgnoreCase(data, search []byte, searchLen int) bool {
	dataLen := len(data)
	for i := 0; i <= dataLen-searchLen; i++ {
		match := true
		for j := 0; j < searchLen; j++ {
			if toLower(data[i+j]) != search[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func regexSearchFile(path string, searchStr string, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
		return
	}
	defer file.Close()

	location := make(map[int]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	regex := regexp.MustCompile("(?i)" + regexp.QuoteMeta(searchStr)) // case-insensitive search

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if regex.MatchString(line) {
			location[lineNum] = line
		}
	}

	if len(location) > 0 {
		resultCh <- Result{path: path, location: location}
	}
}

func boyerMooreSearchFile(path string, searchStr string, resultCh chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s, error: %v\n", path, err)
		return
	}
	defer file.Close()

	location := make(map[int]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	badChar := preprocessBadCharArray(searchStr) // Use fixed-size array
	patternLen := len(searchStr)

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		if boyerMooreMatch(line, searchStr, badChar, patternLen) {
			location[lineNum] = line
		}
	}

	if len(location) > 0 {
		resultCh <- Result{path: path, location: location}
	}
}

func preprocessBadCharArray(pattern string) [256]int {
	var badChar [256]int
	patternLen := len(pattern)

	for i := 0; i < 256; i++ {
		badChar[i] = -1
	}
	for i := 0; i < patternLen; i++ {
		badChar[pattern[i]] = i
	}
	return badChar
}

func boyerMooreMatch(text, pattern string, badChar [256]int, patternLen int) bool {
	textLen := len(text)
	s := 0

	for s <= (textLen - patternLen) {
		j := patternLen - 1

		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}
		if j < 0 {
			return true
		} else {
			s += max(1, j-badChar[text[s+j]])
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func SingleThreadedExample(searchStr string) {
	if verbose {
		defer time_completion.Timer()()
		print("Single Thread -> ")
	}

	resultCh := make(chan Result, len(books))

	// Iterate over each book and perform a direct function call to search
	for _, book := range books {
		basicSearchFile(book, searchStr, resultCh, nil)
	}

	// Collect results from the channel and close it once all are received
	close(resultCh)

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})
}

func basicMultithreadingExample(searchStr string, maxGoroutines int) {
	if verbose {
		print("Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() { <-semaphore }()
			basicSearchFile(book, searchStr, resultCh, &wg)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

}

func regexMultithreadingExample(searchStr string, maxGoroutines int) {
	if verbose {
		print("Regex Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() { <-semaphore }()
			regexSearchFile(book, searchStr, resultCh, &wg)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

}

func basicMultithreadingExampleDisplayResults(searchStr string, maxGoroutines int) {
	if verbose {
		print("Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() { <-semaphore }()
			basicSearchFile(book, searchStr, resultCh, &wg)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

	for _, result := range results {

		lineNumbers := make([]int, 0, len(result.location))
		for lineNum := range result.location {
			lineNumbers = append(lineNumbers, lineNum)
		}
		sort.Ints(lineNumbers)

		if len(lineNumbers) < 25 {
			print(colors.SetColor(fmt.Sprintf("File: %s\n", result.path), tc.Salmon))
			for _, lineNum := range lineNumbers {
				fmt.Printf("Line %d: %s\n", lineNum, result.location[lineNum])
			}
		}
		println()

	}

	total := 0
	for _, result := range results {
		println(filepath.Base(result.path), ":", len(result.location))
		total += len(result.location)
	}
	println("Total: ", total)
}

func boyerMooreMultithreadingExample(searchStr string, maxGoroutines int) {
	if verbose {
		print("Boyer-Moore Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() { <-semaphore }()
			boyerMooreSearchFile(book, searchStr, resultCh, &wg)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

}

func mmapMultithreadingExample(searchStr string, maxGoroutines int) {
	if verbose {
		print("mmap Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() { <-semaphore }()
			mmapSearchFile(book, searchStr, resultCh, &wg)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

}

func mmapBMMultithreadingExample(searchStr string, maxGoroutines int) {
	if verbose {
		print("mmap-BM Multi Thread -> ")
		defer time_completion.Timer()()
	}

	semaphore := make(chan struct{}, maxGoroutines)

	resultCh := make(chan Result, len(books))
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		semaphore <- struct{}{} // Fill a slot in the semaphore

		go func(book string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			mmapBMSearchFile(book, searchStr, resultCh)
		}(book)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})
}

func GatherResults(resultCh chan Result) {

	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].path < results[j].path
	})

	println("Total Results: ", len(results))

	for _, resulting := range results {
		println("Found: ", len(resulting.location))
		for key, _ := range resulting.location {
			println(key)
		}
	}
}

func boyerMooreMatchBM(text, pattern []byte, badChar [256]int) bool {
	textLen, patternLen := len(text), len(pattern)
	s := 0 // Shift

	for s <= (textLen - patternLen) {
		j := patternLen - 1
		for j >= 0 && toLower(text[s+j]) == pattern[j] {
			j--
		}
		if j < 0 {
			return true
		} else {
			badShift := badChar[toLower(text[s+j])]
			s += max(1, j-badShift)
		}
	}
	return false
}

func preprocessBadCharBM(pattern []byte) [256]int {
	var badChar [256]int
	patternLen := len(pattern)

	for i := 0; i < patternLen-1; i++ {
		badChar[pattern[i]] = i + 1 // Use `i + 1` to differentiate unprocessed indices from 0
	}

	return badChar
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
