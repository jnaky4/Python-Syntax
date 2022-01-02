package main

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
const dexnum = iota + 1

func newPokemon(name string) *pokemon{
	p := pokemon{name: name}

	p.dexnum = dexnum
	return &p
}

func (L *List) Insert(newPokemon pokemon){
	list := &Node{
		next: L.head,
		poke: newPokemon,
	}
	if L.head != nil{
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}
func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.poke)
		list = list.next
	}
	fmt.Println()
}

func main(){
	pokedex := List{}
	nameList := [...]string{"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charizard"}

	for k := range nameList {
		pokedex.Insert(*newPokemon(nameList[k]))
	}


	pokedex.Display()
}