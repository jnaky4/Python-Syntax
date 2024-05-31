package main

import (
	"bytes"
	"sync"
)

type SyncedBuffer struct {
	lock    sync.Mutex
	buffer  bytes.Buffer
}

func main() {
	//the zero value for a slice or map type is not the same as an initialized but empty value of the same
	//type. Consequently, taking the address of an empty slice or map composite literal does not have the same effect
	//as allocating a new slice or map value with new.
	p1 := &[]int{}    // p1 points to an initialized, empty slice with value []int{} and length 0
	p2 := new([]int)  // p2 points to an uninitialized slice with value nil and length 0
	println(p1)
	println(p2)
}

/*
Go has 2 allocation primitives
	New and Make

	New
		Built-in function that allocated memory
		Does not initialize the memory
			It only zeros it

		New(T) allocated zeroed storage for a new item of type T and returns its address

Helpful for data structures that that zero value of each type can be used without further initialization

Since the memory returned by new is zeroed, it's helpful to arrange when designing your data structures that
the zero value of each type can be used without further initialization. This means a user of the data structure can
create one with new and get right to work. For example, the documentation for bytes.Buffer states that "the zero value
for Buffer is an empty buffer ready to use." Similarly, sync.Mutex does not have an explicit constructor or Init method.
Instead, the zero value for a sync.Mutex is defined to be an unlocked mutex.
 */
func allocationWithNew(){
	//Values of type SyncedBuffer are also ready to use immediately upon allocation or just declaration.
	//both p and v will work correctly without further arrangement.
	p := new(SyncedBuffer)  // type *SyncedBuffer
	var v SyncedBuffer      // type  SyncedBuffer
	println(p, v)
}

func newAndMakeComparison(){
	var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	var pi *[]int = new([]int)
	*pi = make([]int, 100, 100)

	// Idiomatic:
	x := make([]int, 100)
	println(p, v, x)
}