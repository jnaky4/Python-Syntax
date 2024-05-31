package main


func main(){
	NewDll()
}

type Node struct {
	next, prev *Node
	key   int
	data interface{}
}

//Dll is designed for an LRU cache
type Dll struct{
	head, tail *Node
	len int
}

func NewDll() *Dll{
	return &Dll{}
}

func (l *Dll) Append(key int, value any) {
}

func (l *Dll) Insert(key int, value any) {
}

func (l *Dll) Remove(key int) {

}

func (l *Dll) Find(key int) *Node {
	current := l.head

	for current != nil {
		if current.key == key {
			return current
		}
		current = current.next
	}

	return nil
}
