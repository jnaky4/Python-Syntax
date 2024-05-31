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

	//UTF encodings
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	//UTF8
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

	//padding
	arr := []string{"Taco", "Alabama", "Jordan", "Russia"}
	//pad to left
	fmt.Printf("%10v\n", arr)
	//pad to right
	arrPadded := fmt.Sprintf("%-10v", arr)
	println(arrPadded)

	//specify padding outside of function
	s := -15
	fmt.Printf("%*s%s\n", s, "wow", "interesting")
}


func Printf(format string, v ...interface{}) (n int, err error) {
	return 0, nil
}
