package main

import "fmt"

func main(){

	//empty
	var myGreeting map[string]string
	fmt.Println(myGreeting == nil)

	//make
	var makeMap = make(map[string]string)
	makeMap["Tim"] = "Good morning."
	makeMap["Jenny"] = "Bonjour."

	//shorthand make
	shorthandMakeMap := make(map[string]string)
	shorthandMakeMap["Tim"] = "Good morning."
	shorthandMakeMap["Jenny"] = "Bonjour."

	//composite literal
	clMap := map[string]string{}
	clMap["Tim"] = "Good morning."
	clMap["Jenny"] = "Bonjour."

	//shorthand composite literal
	shCLMap := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	//add or overwrite entry
	shCLMap["Harleen"] = "Howdy"
	//delete
	//deleting empty value doesnt cause errors
	delete(shCLMap, "Tim")
	delete(shCLMap, "DoesntExist")

	//check if value exists
	if val, exists := shCLMap["Harleen"]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	//iterate through map
	for key, val := range shCLMap {
		fmt.Println(key, " - ", val)
	}

	//look at function below to make nested map
	nestMap(3,3,3)
}

func nestMap(i int, j int, k int) map[int]map[int]int{

	firstMap := make(map[int]map[int]int)
	//secondMap := make(map[int]int)
	for a := 0; a < i; a++{
		firstMap[a] = make(map[int]int)
		for b := 0; b < j; b++{
			for c := 0; c < k; c++{
				firstMap[a][b] = c
			}
		}
	}

	fmt.Printf("%v\n", firstMap)
	return firstMap
}