package main

import "awesomeProject/packages/models"

func main(){

	println("Hello")
	u := models.User{
		ID: 2,
		FirstName: "Tricia",
		LastName: "McMillan",
	}
	println(u)

}
