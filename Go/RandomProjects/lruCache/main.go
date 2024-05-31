package main

import "fmt"

func main(){
	dll := Constructor(1)


	//fmt.Printf("Constructor %+v\n", dll)
	dll.Append(1, 1)
	//fmt.Printf("Before %+v\n", dll)
	dll.Append(2,2)
	dll.Append(3,3)
	dll.Append(4,4)
	dll.Append(5,5)
	dll.Append(6,6)
	fmt.Printf("Before %+v\n", dll)
	//fmt.Printf("Head: %+v\n", dll.head)
	//fmt.Printf("Middle: %+v\n", dll.head.next)
	//fmt.Printf("Tail: %+v\n", dll.tail)


	//fmt.Printf("Find 3: %+v\n", dll.Find(3))
	//fmt.Printf("Find 4: %+v\n", dll.Find(4))

	//fmt.Printf("Find 2: %+v\n", dll.Find(2))
	//fmt.Printf("Find 1: %+v\n", dll.Find(1))

	//dll.Remove(5)

	dll.MoveToTail(6)
	fmt.Printf("After %+v\n", dll)
	fmt.Printf("Head: %+v\n", dll.head)
	fmt.Printf("Tail: %+v\n", dll.tail)
}

type Node struct {
	next, prev *Node
	key   int
	value int
}

type LRUCache struct{
	nMap map[int]*Node
	head, tail *Node
	len, capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		nMap: make(map[int]*Node, capacity),
	}
}

//Append sets the tail to the new node, if dll len > capacity, the head will be removed
func (l *LRUCache) Append(key int, value int) {
	nn := &Node{
		key: key,
		value: value,
	}

	l.nMap[key] = nn

	switch l.len {
	case 0:
		l.head = nn
		l.tail = nn
		break
	case 1:
		l.tail = nn
		l.head.next = l.tail
		l.tail.prev = l.head
		break
	default:
		nn.prev = l.tail
		l.tail.next = nn
		l.tail = nn
	}

	l.len++

	if l.len > l.capacity{
		l.Remove(l.head.key)
	}
}

func (l *LRUCache) Remove(key int) {
	switch l.len {
	case 0:
		break
	case 1:
		l.head = nil
		l.tail = nil
		break
	case 2:
		if l.head.key == key{
			l.head = l.tail
		} else {
			l.tail = l.head
		}

		l.head.next = nil
		l.head.prev = nil
		break
	default:
		rn := l.Find(key)
		if rn == nil{ return }

		if l.head.key == key{
			l.head = l.head.next
			if l.head != nil{
				l.head.prev = nil
			}
			break
		}

		if l.tail.key == key{
			l.tail = l.tail.prev
			l.tail.next = nil
			break
		}

		rn.prev.next = rn.next
		rn.next.prev = rn.prev
	}


	delete(l.nMap, key)
	l.len--
}

//Find return nil if not found
func (l *LRUCache) Find(key int) *Node{
	if v, ok := l.nMap[key]; ok{
		return v
	}
	return nil
}

//MoveToTail only works if the value already exists in the Linked List
func (l *LRUCache) MoveToTail(key int) {
	if n := l.Find(key); n != nil{
		l.Remove(key)
		l.Append(n.key, n.value)
	}
}

func (l *LRUCache) Get(key int) int{
	fn := l.Find(key)
	if fn == nil{
		return -1
	}

	l.MoveToTail(key)
	return fn.value
}

func (l *LRUCache) Put(key int, value int){
	fn := l.Find(key)
	if fn == nil{
		l.Append(key, value)
	} else{
		fn.key = key
		fn.value = value
		l.MoveToTail(key)
	}
}
