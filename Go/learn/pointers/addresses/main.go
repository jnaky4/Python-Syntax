package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")

	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	callByValue(x)
	fmt.Println(x) // x is still 5
	callByReference(&x)
	fmt.Println(x) // x is 0

}

func callByReference(z *int) {
	fmt.Printf("%p\n", &z) // address doesnt change
	fmt.Println(z)         // real address
	*z = 0
}
func callByValue(z int) {
	fmt.Printf("%p\n", &z) // address does change
	fmt.Println(z)         // address in func zero
	z = 0
}
