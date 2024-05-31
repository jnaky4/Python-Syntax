package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//str -> int
	s2i()

	//str -> float
	s2f()

	//str -> bool
	//works with strings 1, t, T, true
	s2b()

	//any -> string
	any2s()

	//int -> int16, int32, int64
	i2i(10)
}

func s2i(){
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

func s2f(){
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))
}

func s2b(){
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)
}

func any2s(){
	b := true
	s2 := fmt.Sprintf("%v", b)
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(s2))
}

func i2i(i int){
	fmt.Println(reflect.TypeOf(i))

	i16 := int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32 := int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64 := int64(i)
	fmt.Println(reflect.TypeOf(i64))
}

func f2f(f32 float32){
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	f64 = 1097.655698798798
	fmt.Println(f64)

	f32 = float32(f64)
	fmt.Println(f32)

	//float32 -> 64
	f2f(float32(10.6556))
}

