package main

import (
	"Go/algorithms/boyer_moore_search"
	"fmt"
	"strings"
)

func main(){
	//defer time_completion.Timer()()

	text := "Test There Once was a hidden teSt in the stack named test"
	search := "TEST"

	lower_search := strings.ToLower(search)
	fmt.Printf("%v\n", boyer_moore_search.BuildReadableSkipTable(lower_search))
	fmt.Printf("%v\n", boyer_moore_search.BuildSkipMap(lower_search))
	fmt.Printf("%v\n", boyer_moore_search.BuildSkipMap(search))

	count, locations := boyer_moore_search.Search(text, search)
	fmt.Printf("%v %v\n",count, locations)

}