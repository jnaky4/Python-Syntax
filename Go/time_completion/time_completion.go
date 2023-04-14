package time_completion

import (
	"fmt"
	"github.com/rs/zerolog"
	"reflect"
	"runtime"
	"time"
)

//func main(){
//	defer Timer()()
//	test(20)
//}
//
//func test(i int){
//	for j := 0; j < i; j++{
//		println(i * j)
//	}
//}

func Timer() func(){
	start := time.Now()
	return func() {
		fmt.Printf("elapsed: %s\n", time.Since(start))
	}
}
func FunctionTimer(i interface{}) func(){
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	start := time.Now()
	return func() {
		fmt.Printf("%s function elapsed: %s\n", funcName, time.Since(start))
	}
}
func LogTimer(log *zerolog.Logger, funcInfo string) func (){
	//funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	start := time.Now()
	return func() {
		//log.Info().Msgf("%s function %s elapsed: %s", funcName, time.Since(start))
		log.Info().Msgf("%s time elapsed: %s", funcInfo, time.Since(start))
		//log.Info().Msgf("time elapsed: %s", time.Since(start))
	}
}
