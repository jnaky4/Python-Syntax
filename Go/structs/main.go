package main

func main(){
//	stucts can associate any type of data together
//	values are fixed at compile time

	//defining type user, user type is struc
	type pokemon struct{
		dexnum int
		name string
	}
	//	initialization
	var pikachu pokemon
	pikachu.name = "Pikachu"
	pikachu.dexnum = 25
	println(pikachu.dexnum)

	//multi line initializer need extra comma, expecting curly brace
	raichu := pokemon{
		dexnum: 25,
		name: "Raichu",
	}
	println(raichu)


}

