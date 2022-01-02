package main

import "fmt"

func main(){
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

}