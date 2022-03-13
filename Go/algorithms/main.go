package main

import (
	"Go/algorithms/a_star_search"
	"fmt"
)

func main(){
	//defer time_completion.Timer()()

	//text := "Test There Once was a hidden teSt in the stack named test"
	//search := "TEST"
	//
	//lower_search := strings.ToLower(search)
	//fmt.Printf("%v\n", boyer_moore_search.BuildReadableSkipTable(lower_search))
	//fmt.Printf("%v\n", boyer_moore_search.BuildSkipMap(lower_search))
	//fmt.Printf("%v\n", boyer_moore_search.BuildSkipMap(search))
	//
	//count, locations := boyer_moore_search.Search(text, search)
	//fmt.Printf("%v %v\n",count, locations)

	//fmt.Printf("%v\n", sorting.Insertion_sort([]int {5,4,3,2,1}))
	//arr := a_star_search.GenerateMap(5,10, .2)
	//for i := range arr{
	//	fmt.Printf("%c\n", arr[i])
	//}

	//todo old way

	//g := a_star_search.Graph{}
	//g.GenerateGraph(h,w)
	//
	////Generate Blank map
	//bmap := a_star_search.BlankMap(h,w)
	////Set map indices to match graph
	//for i := range g.Nodes{
	//	bmap[g.Nodes[i].X][g.Nodes[i].Y] = g.Nodes[i].Terrain
	//}

	//todo new way
	h,w := 10,10
	g := a_star_search.Graph{}
	traversalMap := a_star_search.GenerateMap(h,w, .2)
	g.Nodes = a_star_search.MakeGraphFromTraversalMap(traversalMap)
	g.MakeAdjacencyMapFromTraversalMap(h,w)

	d := g.Edges[*g.Nodes[22]]
	fmt.Printf("%d,%d\n", g.Nodes[22].X, g.Nodes[22].Y)
	for i := range d{
		fmt.Printf("%d,%d\n", d[i].X, d[i].Y)
	}






	//for i := range g.Nodes{
	//	fmt.Printf("(%d,%d) %c \n", g.Nodes[i].X, g.Nodes[i].Y, g.Nodes[i].Terrain)
	//}
	//
	//for i := range traversalMap{
	//	fmt.Printf("%c\n", traversalMap[i])
	//}






	//for i := range bmap{
	//	fmt.Printf("%c\n", bmap[i])
	//}

	//g.PrintGraph()


	//d := g.Edges[*g.Nodes[0]]
	//fmt.Printf("%d,%d\n", g.Nodes[0].X, g.Nodes[0].Y)
	//for i := range d{
	//	fmt.Printf("%d,%d\n", d[i].X, d[i].Y)
	//}


	//adjMap := a_star_search.AdjacencyMap(arr)
	//fmt.Printf("%v\n", mapp)
	//fmt.Printf("should be here : %v\n", mapp[2][2])

}

