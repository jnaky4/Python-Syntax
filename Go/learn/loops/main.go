package main

import "fmt"

func main() {
	fourLoopTypes()
	parallelAssignmentLoop()
	breakLoopWithLabel()

	//UTF8
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	//prints key, value
	for k, v := range wellKnownPorts {
		println(k, v)
	}
	//just keys
	for k := range wellKnownPorts {
		println(k)
	}
	//just values
	for _, v := range wellKnownPorts {
		println(v)
	}


}

/*
	Every Loop is a for loop
	4 Types:
		Loop till condition
		Loop till condition with post clause
		Infinite Loops
		Loop Over Collections
*/
func fourLoopTypes(){
	//loop till
	var i int
	for i < 5 {
		println(i)
		i++
	}

	//loop till with post clause, with break
	for i := 1000000; i < 1000100; i++ {

		if i == 1000099 {
			//break loop
			fmt.Printf("Exited Early")
			break
		}
		if i == 1000098 {
			fmt.Printf("SKIP!")
			//skip to next loop
			continue
		}
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	//infinite loop
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}

	//loops over collections
	slice := []int{1, 2, 3}
	//prints index, value
	for i, v := range slice {
		println(i, v)
	}
}

//run multiple variables in a for you should use parallel assignment
func parallelAssignmentLoop() {
	a := []string{"1", "2", "3", "4"}
	// Reverse a

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

//breaking out of a loop with a label
func breakLoopWithLabel(){
	Loop:
		for n := 0; n < 10; n++{
			switch {
			case n == 9:
				break
			case n == 8:
				println("breaking loop")
				break Loop
			}
		}
		println("broke loop")
}