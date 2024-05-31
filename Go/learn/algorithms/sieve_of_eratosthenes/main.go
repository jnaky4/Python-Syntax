package main

import (
	t "Go/time_completion"
	"fmt"
	"math"
	"strconv"
)

func main(){
	//print(strconv.FormatInt(102,4))
	p := []int{2, 1000, 15, 5, 1}

	for i := range p{
		printPrimes(sieve(p[i]), p[i])
	}

}

func sieve(prime int) map[int]bool{
	defer t.Timer()()
	if prime < 2{
		return map[int]bool{}
	}
	if prime == 2{
		return map[int]bool{2: false}
	}

	m := make(map[int]bool, prime)
	m[2] = false
	for i := 3; i <= prime; i++{
		if i % 2 == 0{
			m[i] = true
		} else{
			m[i] = false
		}
	}

	sqr := int(math.Sqrt(float64(prime)))

	for i := 3; i <= sqr; i += 2 {
		if !m[i]{
			m = updateMap(prime, i, m)
		}
	}

	return m
}

func updateMap(limit int, start int, m map[int]bool) map[int]bool{
	counter := start + start

	for counter < limit{
		if !m[counter]{
			m[counter] = true
		}
		counter += start
	}
	return m
}

func printPrimes(a map[int]bool, limit int){
	for i := 2; i <= limit; i++{
		if a[i] == false{
			fmt.Printf("%d, ", i)
		}
	}
	println()
}

/*optimal notes
3: any number can be summed by its digits and if that is divisible by 3 then it is divisble by three
	1234 = 1 + 2 + 3 + 4 = 10 // 3 = 1 no
5: any number with 5 as last digit is divisible by 5

7: 7 21 35 49 63 -> 7 only hits 7, 1, 5, 9, 3 but 5 should be ignored

*/


func isFive(num int) bool{
	defer t.Timer()()
	a := strconv.Itoa(num)
	if a[len(a)-1] == '5'{
		return true
	}
	return false
}