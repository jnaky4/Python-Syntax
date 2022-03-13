package a_star_search

import (
	"Go/time_completion"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	Plain = rune('.')
	Blocker = rune('b')
	Start = rune('s')
	Finish = rune('f')
	Path = rune('X')
)

type Node struct{
	X       int
	Y       int
	Terrain rune
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Terrain)
}

type Graph struct{
	Nodes []*Node
	Edges map[Node][]*Node
	lock  sync.RWMutex
}

func (g *Graph) AddNode(n *Node){
	g.lock.Lock()
	defer g.lock.Unlock()
	g.Nodes = append(g.Nodes, n)
}


func (g *Graph) AddEdges(n1 *Node, n2 *Node){
	g.lock.Lock()
	defer g.lock.Unlock()


	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}
	//d := g.Edges[*n1]


	var found1 bool
	for _, n := range g.Edges[*n1]{
		if n.X == n2.X && n.Y == n2.Y{
			found1 = true
		}
	}
	if !found1{
		g.Edges[*n1] = append(g.Edges[*n1], n2)
	}
	var found2 bool
	for _, n := range g.Edges[*n2]{
		if n.X == n1.X && n.Y == n1.Y{
			found2 = true
		}
	}
	if !found2{
		g.Edges[*n2] = append(g.Edges[*n2], n1)
	}




}

func (g *Graph) PrintGraph() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + " -> "
		near := g.Edges[*g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}

func A_star(maze []int){
	defer time_completion.FunctionTimer(A_star)()

}

//func (g *Graph) GenerateGraph(h int, w int) {
//	for i := 0; i < h; i++{
//		for j := 0; j < w; j++{
//			g.AddNode(g.decideNodeType(h, w, i, j))
//		}
//	}
//	g.generateEdges(h, w)
//}

//func (g *Graph) AddEdges(i int, j int, h int, w int) {
//
//}

//func (g *Graph) decideNodeType(h int, w int, i int, j int) *Node{
//	xstart, ystart  := 0, 0
//	xend, yend := h-1, w-1
//	blockPercent := .2
//
//	if i == xstart && j == ystart{
//		return &Node{X: i, Y: j, Terrain: Start}
//	}
//	if i == xend && j == yend{
//		return &Node{X: i, Y: j, Terrain: Finish}
//	}
//
//	numBlockedTiles := int(float64(h * w) * blockPercent)
//	roll := rand.Intn(h*w)
//	if roll <= numBlockedTiles{
//		return &Node{X: i, Y: j, Terrain: Blocker}
//	}
//
//	return &Node{X: i, Y: j, Terrain: Plain}
//}

func (g *Graph) generateEdges(h int, w int){
	n1 := Node{}
	n2 := Node{}

	for i := 0; i < h; i++{
		for j := 0; j < w; j++{
			for k := range g.Nodes {
				if g.Nodes[k].X == i && g.Nodes[k].Y == j {
					//fmt.Printf("%d, %d found n1\n", i, j)
					n1 = *g.Nodes[k]
				}
			}
			for m := range g.Nodes{
				if i-1 >= 0 {
					if g.Nodes[m].X == i-1 && g.Nodes[m].Y == j{
						//fmt.Printf("%d, %d found n2 top\n", i, j)
						n2 = *g.Nodes[m]
						g.AddEdges(&n1,&n2)
					}
				}
				if i+1 < h{
					if g.Nodes[m].X == i+1 && g.Nodes[m].Y == j{
						n2 = *g.Nodes[m]
						//fmt.Printf("%d, %d found n2 bottom\n", i, j)
						g.AddEdges(&n1,&n2)
					}
				}
				if j-1 >= 0{
					if g.Nodes[m].X == i && g.Nodes[m].Y == j-1{
						n2 = *g.Nodes[m]
						//fmt.Printf("%d, %d found n2 left \n", i, j)
						g.AddEdges(&n1,&n2)
					}
				}
				if j+1 < w {
					if g.Nodes[m].X == i && g.Nodes[m].Y == j+1 {
						n2 = *g.Nodes[m]
						//fmt.Printf("%d, %d found n2 right\n", i, j)
						g.AddEdges(&n1, &n2)
					}
				}
			}
		}
	}
}

func BlankMap(h int, w int) [][]rune{
	//var arr [][]string
	arr := make([][]rune, h)
	for i := 0; i < h; i++{
		arr[i] = make([]rune, w)
	}

	for i := range arr{
		for j := range arr[i]{
			arr[i][j] = Plain
		}
	}
	return arr
}

func GenerateMap(h int, w int, blockPercent float64) [][]rune{

	arr := BlankMap(h, w)

	arr[0][0] = Start
	arr[h-1][w-1] = Finish
	size := h * w

	numBlockedTiles := int(float64(size) * blockPercent)

	r := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(r)

	i := 0
	for i < numBlockedTiles{
		xblock := r1.Intn(h)
		yblock := r1.Intn(w)
		if arr[xblock][yblock] == Plain{
			arr[xblock][yblock] = Blocker
			i++
		}
	}
	return arr
}

func MakeGraphFromTraversalMap(tMap [][]rune) []*Node{
	var nList []*Node

	for i := range tMap{
		for j := range tMap[i]{
			nList = append(nList, &Node{X: i, Y: j, Terrain: tMap[i][j]})
		}
	}
	return nList
}

func (g *Graph) MakeAdjacencyMapFromTraversalMap(h int, w int){
	for i := range g.Nodes{
		n1 := g.Nodes[i]
		//fmt.Printf("n1 : (%d,%d,%c) \n", n1.X, n1.Y, n1.Terrain)
		if n1.X -1 >= 0{
			n2 := g.Nodes[((n1.X-1) * h) + n1.Y]
			g.AddEdges(n1, n2)
			//fmt.Printf("adj : (%d,%d,%c) \n", n2.X, n2.Y, n2.Terrain)
		}
		if n1.X +1 < h{
			n2 := g.Nodes[((n1.X+1) * h) + n1.Y]
			g.AddEdges(n1, n2)
			//fmt.Printf("adj : (%d,%d,%c) \n", n2.X, n2.Y, n2.Terrain)
		}
		if n1.Y -1 >= 0{
			n2 := g.Nodes[((n1.X) * h) + n1.Y-1]
			g.AddEdges(n1, n2)
			//fmt.Printf("adj : (%d,%d,%c) \n", n2.X, n2.Y, n2.Terrain)
		}
		if n1.Y +1 < w{
			n2 := g.Nodes[((n1.X) * h) + n1.Y+1]
			g.AddEdges(n1, n2)
			//fmt.Printf("adj : (%d,%d,%c) \n", n2.X, n2.Y, n2.Terrain)
		}

	}
}


//func AdjacencyMap(traversalMap [][]rune) map[int]map[int]map[int]map[int]int{
//	adjacenyMap := make(map[int]map[int]map[int]map[int]int)
//
//	for i := range traversalMap{
//		adjacenyMap[i] = make(map[int]map[int]map[int]int)
//		for j := range traversalMap[i]{
//			adjacenyMap[i][j] = make(map[int]map[int]int)
//			adjacenyMap[i][j] =	adjacency(i, j, traversalMap)
//		}
//	}
//	return adjacenyMap
//}
//
//func adjacency(i int, j int, traversalMap [][]rune) map[int]map[int]int{
//	if traversalMap[i][j] == Blocker{
//		return nil
//	}
//
//	adjacent := make(map[int]map[int]int)
//	adjacent[i] = make(map[int]int)
//	if i-1 >= 0 {
//		adjacent[i-1] = make(map[int]int)
//	}
//	if i+1 < len(traversalMap){
//		adjacent[i+1] = make(map[int]int)
//	}
//
//	if i-1 >= 0 && traversalMap[i-1][j] != Blocker{
//		adjacent[i-1][j] = j
//	}
//	if i+1 < len(traversalMap) && traversalMap[i+1][j] != Blocker{
//		adjacent[i+1][j] = j
//	}
//	if j-1 >= 0 && traversalMap[i][j-1] != Blocker{
//		adjacent[i][j-1] = j-1
//	}
//	if j+1 < len(traversalMap[0]) && traversalMap[i][j+1] != Blocker{
//		adjacent[i][j+1] = j+1
//	}
//	//fmt.Printf("%v\n", adjacent)
//	return adjacent
//
//}

