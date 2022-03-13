package main

import "fmt"

func main(){
	//----Arrays----
	//	fixed sized collection of similar data types
	//	long syntax
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1:])

	//	implicit syntax
	arr2 := [3] int{1,2,3}
	fmt.Println(arr2)

	//----Slices----

	//A slice has both a length and a capacity.
	//length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array,
	//counting from the first element in the slice.

	//declare empty slice
	var intSlice []int
	fmt.Printf("%*s \tLen: %-2v \tCap: %-2v\n", -11, "intSlice", len(intSlice), cap(intSlice))
	//declare slice literal
	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Printf("%*s \tLen: %-2v \tCap: %-2v \t%v\n", -11, "strSlice", len(strSlice), cap(strSlice), strSlice)
	//declare with new keyword
	var newIntSlice = new([50]int)[0:10]
	fmt.Printf("%*s \tLen: %-2v \tCap: %-2v\n", -11, "newIntSlice", len(newIntSlice), cap(newIntSlice))
	//declare with make
	var makeIntSlice = make([]int, 10)
	fmt.Printf("%*s \tLen: %-2v \tCap: %-2v\n", -11, "makeIntSlice", len(makeIntSlice), cap(makeIntSlice))
	var makeStrSlice = make([]string, 10, 20) // when length and capacity is different
	fmt.Printf("%*s \tLen: %-2v \tCap: %-2v\n", -11, "makeStrSlice", len(makeStrSlice), cap(makeStrSlice))






	//sliceExamples()
	//mapExample()

}
//	slice
//	built on top of array
func sliceExamples() {
	arr := [3]int{1,2,3}
	slice := arr[:]
	//any changes to slice/array wil be reflected in the other
	slice[1] = 2
	fmt.Println(slice, arr)

	//slice is not a fixed sized entity
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2)

	//start a 1 to end
	s2 := slice2[1:]
	//up to but not 2
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(s2, s3, s4)

}
func mapExample(){
	//map string datatype to an integer
	m := map[string]int{"Pikachu":25}
	fmt.Println(m["Pikachu"])

	//add to map
	m["Raichu"] = 26

	//delete from map
	delete(m, "Raichu")

}