package main

import "fmt"

const p = "death & taxes"

func main() {
	//Println
	fmt.Println("Hello world!")
	fmt.Println(42)
	//Printf
	fmt.Printf("%d\n", 42)

	const q = 42
	fmt.Println("q - ", q)
	//above main
	fmt.Println("p - ", p)

	//binary
	fmt.Printf("%d - %b \n", 42, 42)

	//hexadecimal
	fmt.Printf("%d - %x \n", 42, 42)
	fmt.Printf("%d - %#x \n", 42, 42)
	fmt.Printf("%d - %#X \n", 42, 42)

	//memory address
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)


	//padding
	arr := [] string{"Taco", "Alabama", "Jordan", "Russia"}
	//pad to left
	fmt.Printf("%10v\n", arr)
	//pad to right
	arrPadded := fmt.Sprintf("%-10v", arr)
	println(arrPadded)

	//specify padding outside of function
	s := -15
	fmt.Printf("%*s%s\n", s,"wow", "interesting")
}
