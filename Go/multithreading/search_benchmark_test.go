package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
)

var testPath = "your_test_file.txt" // Path to a file to be used in the benchmark
var searchStr = "time"              // The string to search for

//go test -bench=.

func stringSearchFile(path, searchStr string) (map[int]string, error) {
	location := make(map[int]string)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), searchStr) {
			location[lineNum] = line
		}
	}
	return location, scanner.Err()
}

func regexSearchFile(path, searchStr string) (map[int]string, error) {
	location := make(map[int]string)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	re := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(searchStr))
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if re.MatchString(line) {
			location[lineNum] = line
		}
	}
	return location, scanner.Err()
}

// Benchmark test for the `stringSearchFile` function
func BenchmarkStringContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringSearchFile(testPath, searchStr)
	}
}

// Benchmark test for the `regexSearchFile` function
func BenchmarkRegexMatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		regexSearchFile(testPath, searchStr)
	}
}
