package algos

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Simple realisation of A linked list data structure

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
	if list.head == nil {
		list.head = node
	}
}

// DeleteTail deletes last node and move the pointer to another element
func (list *List) DeleteTail() {
	if !isEmpty(*list) {
		node := list.tail
		list.tail = node.prev
		node.prev.next = nil
		node = nil
		list.size--
		log.Println("Tail deleted")
	}
}

// DeleteHead deletes first node and move head to the next element
func (list *List) DeleteHead() {
	if !isEmpty(*list) {
		node := list.head
		list.head = node.next
		node = nil
		list.size--
		log.Println("Head deleted")
	}
}

// AddNodeToPosition adds node specified positionx
func (list *List) AddNodeToPosition(position, value int) bool {
	if list.isBadPosition(position) {
		return false
	}
	node := list.head
	var current int
	for true {
		if position == current {
			newNode := new(Node)
			newNode.value = value
			if node.next != nil {
				nextNode := node.next
				node.next = newNode
				newNode.prev = node
				newNode.next = nextNode
				nextNode.prev = newNode
			} else {
				node.next = newNode
				newNode.prev = node
			}
			list.size++
			return true
		}
		node = node.next
		current++
	}
	return false
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
	if list.isBadPosition(position) {
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

// RemoveElementByPosition from position
func (list *List) RemoveElementByPosition(position int) bool {
	if list.isBadPosition(position) {
		return false
	}
	if position == 0 {
		node := list.head
		list.head = node.next
		node = nil
		list.size--
		return true
	}
	if position == list.size-1 {
		node := list.tail
		list.tail = node.prev
		node.prev.next = nil
		node = nil
		list.size--
		return true
	}
	node := list.head
	var counter int
	for node != nil {
		if position == counter {
			fmt.Println("Removing Node with value: ", node.value)
			node.prev.next = node.next
			node.next.prev = node.prev
			node = nil
			list.size--
			return true
		}
		node = node.next
		counter++
	}
	return false
}

// RemoveElementByValue removing first element, return true if element is found and removed
func (list *List) RemoveElementByValue(value int) bool {
	if list.size == 1 && list.head.value == value {
		list.head = nil
		list.tail = nil
		list.size--
		return true
	}
	if list.head.value == value {
		node := list.head
		list.head = node.next
		node = nil
		list.size--
		return true
	}
	if list.tail.value == value {
		node := list.tail
		list.tail = node.prev
		list.tail.next = nil
		node = nil
		list.size--
		return true
	}
	node := list.head.next
	for node != nil {
		if node.value == value {
			node.prev.next = node.next
			node.next.prev = node.prev
			node = nil
			list.size--
			return true
		}
		node = node.next
	}

	log.Println("Such element didnt exists")
	return false
}

// isEmpty func checks if list is empty
func isEmpty(list List) bool {
	return list.size == 0
}

func (list List) isBadPosition(position int) bool {
	if position >= list.size || position < 0 {
		log.Println("Bad position!")
		return true
	}
	return false
}

// IterList iter through the list
func IterList(list List) {
	if isEmpty(list) {
		fmt.Println("List is empty")
		return
	}
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

// WriteToFile function writes our node values and some information to a specified file
// returns an error, if file path is bad
func (list List) WriteToFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("File not exists!")
		return err
	}
	file, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()
	bytesWritten, err := file.Write([]byte(getStringReprOfList(list)))
	checkErr(err)
	log.Printf("Wrote %d bytes.\n", bytesWritten)
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Println("Problems with file")
	}
}

func getStringReprOfList(list List) string {
	var repr string
	repr = "List name: " + list.name + "\n"
	repr += "List size: " + strconv.Itoa(list.size) + "\n"
	repr += "List Node values:\n\n"
	node := list.head
	for node != nil {
		repr += "Value: " + strconv.Itoa(node.value) + "\n"
		node = node.next
	}
	return repr
}

// GetSize return a list size
func (list List) GetSize() int {
	return list.size
}

// GetListName returns a list name
func (list List) GetListName() string {
	return list.name
}

// GetHead return head of a list
func (list List) GetHead() Node {
	return *list.head
}

// GetNodeValue returns node value
func (node Node) GetNodeValue() int {
	return node.value
}

// GetNextNode returns next node
func (node Node) GetNextNode() Node {
	return *node.next
}

// RunList function
func RunList() {
	list := GetList("It`s a kkinda magic")
	list.AddNode(13)
	list.AddNode(100)
	list.AddNode(755)
	list.WriteToFile("algos/file.txt")
	//IterList(*list)
}
