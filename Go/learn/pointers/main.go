package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Println(*b)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

	//	holds address that points to value
	//* pointer operator
	var firstName *string = new(string)
	//* dereference operator
	*firstName = "Jacob"
	//prints address
	fmt.Println(firstName)
	//prints value
	fmt.Println(*firstName)

	lastName := "Alongi"
	fmt.Println(lastName)
	ptr := &lastName
	fmt.Println(ptr)
	lastName = "Rhaman"
	fmt.Println(ptr, &ptr)

	/* local variable definition */
	var c int = 100
	var d int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap(&c, &d)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put csv into y */
}
