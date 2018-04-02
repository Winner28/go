package algos

import "fmt"

// Node represents a simple node
type Node struct {
	next  *Node
	prev  *Node
	value int
}

// List repr List
type List struct {
	node *Node
	head *Node
	tail *Node
	size int
	name string
}

// AddNode method adds node to a list
func (list *List) AddNode(value int) {
	node := new(Node)
	node.value = value
	node.next = nil
	node.prev = list.node
	list.node.next = node
	list.node = node
	list.tail = list.node
	list.size++
}

// IterList iter through the list
func (list *List) IterList() {
	node := list.head
	fmt.Println("List size: ", list.size)
	fmt.Println("Head value: ", list.head.value)
	fmt.Println("Tail value: ", list.tail.value)
	fmt.Println()
	for node != nil {
		fmt.Println("Node value: ", node.value)
		node = node.next
	}
}

// Init initialization
func Init(name string) *List {
	list := new(List)
	list.name = name
	return list.initList()
}

func (list *List) initList() *List {
	node := new(Node)
	node.value = 228
	node.next = nil
	node.prev = nil
	list.node = node
	list.head = list.node
	list.tail = list.head
	list.size = 1
	return list
}

// RunList function
func RunList() {
	list := Init("It`s a kkinda magic")
	list.AddNode(50)
	list.AddNode(100)
	list.AddNode(755)
	list.IterList()
}
