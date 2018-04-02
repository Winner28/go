package algos

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
}

//AddNode method ads Node
func (node *Node) AddNode(value int) {

	newNode := Node{
		next:  nil,
		prev:  node,
		value: value,
	}
	node.next = &newNode
}

// AddNode method adds node to a list
func (list *List) AddNode(value int) {

}

// FirstInitialization first
func FirstInitialization() *List {
	list := new(List)
	return list.initList()
}

func (list *List) initList() *List {
	node := new(Node)
	node.value = 100
	node.next = nil
	node.prev = nil
	list.node = node
	list.head = list.node
	list.tail = list.head
	return list
}
