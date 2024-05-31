package main

import "fmt"

func main() {
	//println(isEven(1))
	//println(removeLastBit(5))

	//println(addBit(11))
	//println(addNBits(5, 3))
	//println(removeNBits(255, 7))

	fmt.Printf("set bit (%d) %b -> (%d) %b\n", 8, 8, setBit(8, 0), setBit(8, 0))

	fmt.Printf("clear bit (%d) %b -> (%d) %b\n", 7, 7, clearBit(7, 1), clearBit(7, 1))
}

func isEven(n int) bool {
	//and last bit with 1
	//111 & 001 = 1
	//110 & 1 = 0
	return n&1 == 0
}

// removing a bit is the right shift operator
func removeLastBit(n int) int {
	//5 >> 1 = 2
	//101 >> 1 = 10

	//right shift is the same as truncated division by 2
	// 5 >> 1 == 5/2

	return n >> 1
}

// add bit is the left shift operator
func addBit(n int) int {
	//5 << 1
	//101 << 1 = 1010 (10)
	//left shift is the same as multiply by 2

	return n << 1
}

// adding n bits is left shift operator n times
func addNBits(x int, n int) int {
	//5 << 2
	//101 << 2 = 10100 (20)
	return x << n
}

func removeNBits(x int, n int) int {
	return x >> n
}

// A^A = 0, A^B^B = A, A^B^B^C^C = A
// xor the output is true , only if both the inputs are of opposite kind
func xor(x int, y int) int {
	return x ^ y
}

func and(x int, y int) int {
	return x & y
}

func or(x int, y int) int {
	return x | y
}

func not(x int) int {
	return ^x
}

func andNot(x int, y int) int {
	return x &^ y
}

// setBit shifts 1 left i times, i being the index to set and then or's the value to set the bit
func setBit(x int, i int) int {
	//set (8) 1000, index 1 setBit(8,1)
	// 1 << 1 -> 10
	// 1000 | 0010 -> (10) 1010

	return x | 1<<i
}

// clearBit shifts a bit i times, then nots the value, then ands with the number
func clearBit(x int, i int) int {

	//1 << 1 = 10, ^10 -> 01 or 11111101
	//but really is this
	//00000111 and 11111101 = 00000101

	return x & ^(1 << i)
}
func toggleBit(x int, i int) {}
