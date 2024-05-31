package time_completion

import (
	"fmt"
	"github.com/rs/zerolog"
	"reflect"
	"runtime"
	"sync"
	"time"
)

//func main(){
//	defer Timer()()
//	test(20)
//	PrintTimerStatistics()
//}
//
//func test(i int){
//	defer FunctionTimerCounter(test)()
//	for j := 0; j < i; j++{
//		println(i * j)
//	}
//}

// FunctionStats represents the statistics for a function.
type FunctionStats struct {
	TotalTime time.Duration
	Count     int
}

var (
	statsMap   = make(map[string]FunctionStats)
	statsMutex sync.Mutex
)

// FunctionTimerCounter measures the time taken for a function to complete and stores the statistics.
func FunctionTimerCounter(i interface{}) func() {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	start := time.Now()

	return func() {
		elapsed := time.Since(start)

		statsMutex.Lock()
		defer statsMutex.Unlock()

		stats, exists := statsMap[funcName]
		if !exists {
			stats = FunctionStats{}
		}

		stats.TotalTime += elapsed
		stats.Count++

		statsMap[funcName] = stats

		fmt.Printf("%s function elapsed: %s\n", funcName, elapsed)
	}
}

// PrintTimerStatistics prints the measurements of each function that called FunctionTimerCounter
func PrintTimerStatistics() {
	statsMutex.Lock()
	defer statsMutex.Unlock()

	for funcName, stats := range statsMap {
		averageTime := stats.TotalTime / time.Duration(stats.Count)
		fmt.Printf("%s - Total Time: %s, Average Time: %s\n", funcName, stats.TotalTime, averageTime)
		fmt.Printf("%s - Times Function Executed: %d", funcName, stats.Count)
	}

}

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("elapsed: %s\n", time.Since(start))
	}
}

func FunctionTimer(i interface{}) func() {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	start := time.Now()
	return func() {
		fmt.Printf("%s function elapsed: %s\n", funcName, time.Since(start))
	}
}

func LogTimer(log *zerolog.Logger, funcInfo string) func() {
	//funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Id()
	start := time.Now()
	return func() {
		//log.Info().Msgf("%s time elapes: %s", funcName, time.Since(start))
		log.Info().Msgf("%s time elapsed: %s", funcInfo, time.Since(start))
	}
}
