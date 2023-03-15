package time_completion

import(
	"fmt"
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
