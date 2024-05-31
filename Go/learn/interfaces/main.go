package main

import (
	"fmt"
	"math"
	"sort"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

// which implements the shape interface
func (s square) area() float64 {
	return s.side * s.side
}

// another shape
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
	fmt.Println("Total Area: ", totalArea(c, s))
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

/*
	A struct in go can use other methods for other packages if it implements the required interface methods

	example
	If you implement the sort.Interface - Len(), Less(i, j int) bool, and Swap(i, j int) you can apply sort.Sort()
	to your struct
 */
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}