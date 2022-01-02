package main

import "fmt"


var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func main() {
	fmt.Println(product, quantity, price, inStock)

	//cannot reassign, have to initialize during declaration
	const (
		//implicitly typed constant
		pi = 3.14
		language = "Go"
	)
	fmt.Println(pi,language)

	//Declaration
	var message string
	message = "Hello World."
	fmt.Println(message)

	//Null Values
	var aa int
	var bb string
	var cc float64
	var dd bool

	fmt.Println()
	fmt.Println("Null Variables")
	fmt.Printf("%v \n", aa)
	fmt.Printf("%v \n", bb)
	fmt.Printf("%v \n", cc)
	fmt.Printf("%v \n", dd)

	fmt.Println()

	//many at once
	//var a, b, c = 1, 2, 3
	//var a, b, c int
	//var a, b, c int = 1, 2, 3
	//var message = "Hello World!"

	var flo float64 = 3.14
	var floa float32 = 3.14
	fmt.Printf("%v\nb", flo)
	fmt.Printf("%v \n", floa)

	//implicit initialization syntax
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	//variables
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	//types
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)

}
