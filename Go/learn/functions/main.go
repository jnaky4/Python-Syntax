package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//anonymous self executing function
	func() {
		fmt.Println("I'm driving!")
	}()

	port := 3000
	port, err := startWebServer(port, 2)
	println(port, err)

	//could just make the param type array instead of using a splat to pass in one at a time
	data := []string{"red", "blue", "green", "yellow"}
	variadicExample(data...)

	variadicExample2(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})

	deferExample()
	reverseWithDefer()
	deferTracing()
}

// if matching param types only need to specify one
// returns int and error
func startWebServer(port int, retries int) (int, error) {
	println("Starting server...")

	println("Server started on port", port)
	println("Number of retries", retries)
	if port > 9999 {
		return port, errors.New("Bad port")
	}
	return port, nil
}

// variadic function
func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

// accepts an arbitrary number of arguments of type slice.
func variadicExample2(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

// A defer statement is often used with paired operations like open and close,
// connect and disconnect, or lock and unlock to ensure that resources are released
// in all cases, no matter how complex the control flow. The right place for a defer
// statement that releases a resource is immediately after the resource has been
// successfully acquired.
func deferExample() {
	defer println("Called once function closes")
	println("Filler")
}

//defer functions are executed in LIFO order
func reverseWithDefer(){
	a := []int{1,2,3,4,5,6,7,8,9}
	for _, i := range a{
		defer println(i)
	}
}

//by exploiting the fact that arguments to deferred functions are evaluated when the defer executes.
//The tracing routine can set up the argument to the untracing routine
func deferTracing(){
	defer un(trace("deferTracing"))
	fmt.Println("in deferTracing")
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}