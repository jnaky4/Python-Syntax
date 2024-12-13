package streaming

import (
	"Go/time_completion"
	"bufio"
	"net/http"
	"os"
)

//func StreamBookHandler(w http.ResponseWriter, r *http.Request) {
//	defer time_completion.Timer()()
//	bookPath := "/Users/Z004X7X/Git/syntax/Go/multithreading/books/theLordOfTheRingsTrilogy.txt"
//	bookFile, err := os.Open(bookPath)
//	if err != nil {
//		http.Error(w, "Unable to open book file", http.StatusInternalServerError)
//		return
//	}
//	defer bookFile.Close()
//
//	w.Header().Set("Content-Type", "text/plain")
//	w.WriteHeader(http.StatusOK)
//
//	writer := bufio.NewWriter(w)
//	lines := make(chan string, 100)
//
//	go func() {
//		scanner := bufio.NewScanner(bookFile)
//		for scanner.Scan() {
//			lines <- scanner.Text()
//		}
//		close(lines)
//	}()
//
//	for line := range lines {
//		_, err := writer.WriteString(line + "\n")
//		if err != nil {
//			break
//		}
//		writer.Flush()
//	}
//}

func StreamBookHandler(w http.ResponseWriter, r *http.Request) {
	defer time_completion.Timer()()
	bookPath := "/Users/Z004X7X/Git/syntax/Go/multithreading/books/theLordOfTheRingsTrilogy.txt"
	bookFile, err := os.Open(bookPath)
	if err != nil {
		http.Error(w, "Unable to open book file", http.StatusInternalServerError)
		return
	}
	defer bookFile.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	writer := bufio.NewWriterSize(w, 32*1024) // Use a 32KB buffer for better performance
	scanner := bufio.NewScanner(bookFile)

	for scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			return
		}
	}

	writer.Flush()

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
	}
}

//func StreamBookHandler(w http.ResponseWriter, r *http.Request) {
//	defer time_completion.Timer()()
//	bookPath := "/Users/Z004X7X/Git/syntax/Go/multithreading/books/theLordOfTheRingsTrilogy.txt"
//	bookFile, err := os.Open(bookPath)
//	if err != nil {
//		http.Error(w, "Unable to open book file", http.StatusInternalServerError)
//		return
//	}
//	defer bookFile.Close()
//
//	w.Header().Set("Content-Type", "text/plain")
//	w.WriteHeader(http.StatusOK)
//
//	// Use a larger buffer size for the writer
//	writer := bufio.NewWriterSize(w, 64*1024) // Increase buffer to 64KB for more efficient writing
//
//	// Use a larger buffer for the scanner to avoid reallocation
//	scanner := bufio.NewScanner(bookFile)
//	scanner.Buffer(make([]byte, 32*1024), 32*1024) // Scanner buffer set to 64KB
//
//	var line string
//	for scanner.Scan() {
//		line = scanner.Text() // Scan the next line
//
//		// Write the line into the writer buffer
//		_, err := writer.WriteString(line + "\n")
//		if err != nil {
//			http.Error(w, "Error writing data", http.StatusInternalServerError)
//			return
//		}
//	}
//
//	// Flush the writer buffer once all lines are processed
//	writer.Flush()
//
//	if err := scanner.Err(); err != nil {
//		http.Error(w, "Error reading file", http.StatusInternalServerError)
//	}
//}
