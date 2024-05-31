package main

import (
	"fmt"
)

func main() {
	println(sumIndicesWithKSetBits([]int{5,10,1,5,2}, 1))
}


//2859 TODO look at optimal solution
func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	t := ""
	for i := 0; i < len(nums); i++{
		t = fmt.Sprintf("%b", i)
		c := 0
		for j := 0; j < len(t); j++{
			if t[j] == 49{
				c++
			}
		}
		if c == k{
			sum += nums[i]
		}
	}

	return sum
}
