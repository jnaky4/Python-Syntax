package main

import "fmt"

type Node struct {
	value int
	visited bool
	edges []*Node
}
func DFS(n *Node){
	if n == nil{ return }
	n.visited = true

	fmt.Printf("%d ", n.value)

	for _, edge := range n.edges {
		if !edge.visited {
			DFS(edge)
		}
	}
}

func main(){
	node1 := &Node{value: 10}
	node2 := &Node{value: 20}
	node3 := &Node{value: 30}
	node4 := &Node{value: 40}
	node5 := &Node{value: 50}

	node1.edges = []*Node{node2, node3}
	node2.edges = []*Node{node4, node5}
	node3.edges = []*Node{node5}

	DFS(node1)
}
