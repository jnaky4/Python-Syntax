package main

import "fmt"

// In general, you should use enums whenever you have a group of limited, and related values for a particular type.

// Direction Idiomatic way of implementing enumerated types
type Direction int

const (
	North Direction = iota
	East
	South
	West
	//NorthEast throws error
)

// replaces return values of iota (0,1,2...) to North, East, South, West
// will no longer recognize new Enums unless specified here
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("iotas()")
	iotas()

	var n, e, s, w = North, East, South, West
	fmt.Println(n, e, s, w)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

}

func iotas() {
	const (
		// Iotas, represents successive integer constants
		// Starts at zero, skips 0 starts at 1
		first = iota + 1
		//will use same iota and increment as well
		second
		third
		//skips
		_
	)
	// iota resets at new constant blocks
	const (
		fourth = iota
		fifth
	)

	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth)

}
