package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)


//confused what is actually happening to byte values
func main(){
	b := []byte{48,50,57}
	x := 0
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool{
	r, _ := utf8.DecodeRune([]byte{b})
	println("r: ", r)
	if unicode.IsDigit(r){
		return true
	}
	return false
}

