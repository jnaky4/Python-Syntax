package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

//func main() {
//	fmt.Printf("%d\n", time_completion.TimeFunction(func() { time_completion.FormatDuration(1) }))
//	fmt.Printf("%d\n", time_completion.TimeFunction(func() { time_completion.FormatDuration2(1) }))
//
//}

//func main() {
//	//println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
//	// Get the number of CPUs available to the Go scheduler
//
//}

// 2859 TODO look at optimal solution
func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	t := ""
	for i := 0; i < len(nums); i++ {
		t = fmt.Sprintf("%b", i)
		c := 0
		for j := 0; j < len(t); j++ {
			if t[j] == 49 {
				c++
			}
		}
		if c == k {
			sum += nums[i]
		}
	}

	return sum
}

func cleanPokemonFiles() {
	getwd, _ := os.Getwd()
	dir := path.Join(getwd, "images", "pokemon")
	files := ListDirFiles(dir)

	//Removing files with end Shiny or 2 or 3
	//for _, v := range files{
	//	if v[len(v)-5:len(v)-4] == "2"{
	//		os.Remove(v)
	//	}
	//	if v[len(v)-5:len(v)-4] == "3"{
	//		os.Remove(v)
	//	}
	//	if v[len(v)-9:len(v)-4] == "Shiny"{
	//		os.Remove(v)
	//	}
	//}

	//Renaming to dexnum.png
	for _, v := range files {
		u := strings.Split(v, "/")
		r, _ := strconv.Atoi(u[len(u)-1][:3])
		if r == 0 {
			continue
		}
		_ = os.Rename(v, fmt.Sprintf("%s/%s.png", dir, strconv.Itoa(r)))
	}

}

func ListDirFiles(filepath string) []string {
	var files []string
	dir, err := os.ReadDir(filepath)
	if err != nil {
		return nil
	}
	for _, file := range dir {
		//info, err := file.Info()
		//if err != nil {
		//	return nil
		//}
		//fmt.Printf("%+v\n", info)
		files = append(files, path.Join(filepath, file.Name()))
	}
	return files
}
