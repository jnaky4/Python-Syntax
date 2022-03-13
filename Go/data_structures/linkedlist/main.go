package main

//todo add lock https://flaviocopes.com/golang-data-structure-linked-list/


import (
	//"awesomeProject/structs/pokemon"
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	poke pokemon
}

type List struct {
	head *Node
	tail *Node

}
type pokemon struct{
	name string
	dexnum int

}

func newPokemon(name string, dexnum int) *pokemon{
	p := pokemon{name: name, dexnum: dexnum}
	return &p
}
//todo add insert
//todo add delete
//todo add lock

func (L *List) Prepend(newPokemon pokemon){
	newNode := &Node{
		next: L.head,
		poke: newPokemon,
	}

	if L.head != nil{
		L.head.prev = newNode

		if L.tail == nil{
			L.tail = L.head.next
		}
	}


	L.head = newNode
}

func (L *List) Append(newPokemon pokemon){
	newNode := &Node{
		prev: L.tail,
		poke: newPokemon,
	}

	if L.head == nil{
		L.head = newNode
		return
	}

	if L.tail == nil {
		L.head.next = newNode
	} else{
		L.tail.next = newNode
	}

	L.tail = newNode
}

func (l *List) Display() {
	curNode := l.head
	for curNode != nil {
		if curNode.next != nil{
			fmt.Printf("%+v ->", curNode.poke)
		} else {
			fmt.Printf("%+v\n", curNode.poke)
		}

		curNode = curNode.next
	}
	fmt.Println()
}

func main(){
	pokedex := List{}
	nameList := [...]string{"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon", "Charizard"}
	for j := range nameList{
		pokedex.Prepend(*newPokemon(nameList[j], j + 1))
	}

	pokedex.Display()
	fmt.Printf("head: %+v\n", pokedex.head.poke)
	fmt.Printf("tail: %+v\n", pokedex.tail.poke)

	for k := range nameList {
		pokedex.Append(*newPokemon(nameList[k], len(nameList) + k+1))
	}

	fmt.Printf("head: %+v\n", pokedex.head.poke)
	fmt.Printf("tail: %+v\n", pokedex.tail.poke)

	pokedex.Display()
}