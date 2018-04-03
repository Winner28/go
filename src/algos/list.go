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
	node        *Node
	head        *Node
	tail        *Node
	size        int
	name        string
	initialized bool
}

var globalList *List

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

// DeleteTail deletes last node and move the pointer to another element
func (list *List) DeleteTail() {

}

// DeleteHead deletes first node and move head to the next element
func (list *List) DeleteHead() {

}

// AddNodeToPosition adds node specified positionx
func (list *List) AddNodeToPosition(position, value int) {

}

// Contains element
func (list List) Contains(value int) bool {
	node := list.head
	var counter int
	for node != nil {
		if node.value == value {
			fmt.Printf("We found value on position: %v\n", counter)
			return true
		}
		node = node.next
		counter++
	}
	return false
}

// GetElement from position
func (list List) GetElement(position int) (int, bool) {
	if position > list.size || position < 0 {
		fmt.Println("Bad position!")
		return -1, false
	}
	var counter int
	if position > list.size/2 {
		fmt.Println("Starts from an end")
		node := list.tail
		counter = list.size
		for counter > position {
			node = node.prev
			counter--
		}
		fmt.Println("We got this value: ", node.value)
		return node.value, true
	}
	node := list.head
	for counter < position {
		node = node.next
		counter++
	}
	fmt.Println("We got this value: ", node.value)
	return node.value, true

}

// RemoveElement from position
func (list *List) RemoveElement(position int) bool {
	return false
}

// IterList iter through the list
func IterList(list List) {
	node := list.head
	fmt.Println(list.name, " list")
	fmt.Println("List size: ", list.size)
	fmt.Println("Head value: ", list.head.value)
	fmt.Println("Tail value: ", list.tail.value)
	fmt.Println()
	for node != nil {
		fmt.Println("Node value: ", node.value)
		node = node.next
	}
}

// GetList initialization
func GetList(name string) *List {
	if globalList != nil {
		fmt.Println("We already got the list!")
		return globalList
	}
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
	list.initialized = true
	globalList = list
	return list
}

// RunList function
func RunList() {
	list := GetList("It`s a kkinda magic")
	list.AddNode(50)
	list.AddNode(100)
	list.AddNode(755)
	list.Contains(755)
	//IterList(*list)
}
