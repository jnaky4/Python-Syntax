package main

/*
	function literal
		A Function literal is a function that is not declared but that is passed in as an expression.
		Lambdas and anonymous functions ARE function literals.
	function declaration
		binds an identifier ie the function name to a function
	signature
		the parameters and returned values to an expression.
		func expression(param int) returned int{}
	expression
		An expression specifies the computation of a value by applying operators and functions to operands
	statement
	composite literal
	operands
	receiver



*/

func main(){
	compositeLiteral()
	functionLiteral()
}

type example struct{
	name string
	age int
	omitted string
}
//composite literal is an expression that creates a new instance each time it is evaluated.
func compositeLiteral(){
	//slice composite literal
	cl := []int{23, 56, 89, 34}
	println(cl)
	cle := example{name: "jake", age: 31}
	println(cle)
	//Taking the address of a composite literal generates a pointer to a unique
	//variable initialized with the literal's value.
	var PointExample = &cle
	println(PointExample)
}


//Note that, unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with
//the variable survives after the function returns. In fact, taking the address of a composite literal allocates a
//fresh instance each time it is evaluated, so we can combine these last two lines.
//return &File{fd, name, nil, 0}

//The fields of a composite literal are laid out in order and must all be present. However, by labeling the elements
//explicitly as field:value pairs, the initializers can appear in any order, with the missing ones left as their
//respective zero values. Thus we could say return &File{fd: fd, name: name}

//
func expression(){

}

//	function literal
//		A Function literal is a function that is not declared but that is passed in as an expression.
//		Lambdas and anonymous functions ARE function literals.
func functionLiteral(){
	func(){ println("hello") }()
}

