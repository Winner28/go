package tests

import (
	"algos"
	"testing"
)

var listName string = "My List!"

//TestQuickSort is working well
func TestList(t *testing.T) {
	testList := algos.GetList(listName)
	if !assertEquals(testList.GetListName(), listName) {
		t.Error("Names not equals!")
	}
	if !assertEquals(testList.GetSize(), 1) {
		t.Error("Sizes are not equals!")
	}
	testList.DeleteHead()
	if !assertEquals(testList.GetSize(), 0) {
		t.Error("Error after deletion")
	}
	values := []int{1, 2, 3, 4, 5}
	for _, value := range values {
		testList.AddNode(value)
	}
	if !assertEquals(testList.GetSize(), len(values)) {
		t.Error("Size is different!")
	}
	node := testList.GetHead()
	for i := 0; i < len(values)-1; i++ {
		if !assertEquals(values[i], node.GetNodeValue()) {
			t.Error("Insertion into a list is bad!!!", values[i], node.GetNodeValue())
		}
		node = node.GetNextNode()
	}
	if !testList.Contains(5) {
		t.Error("Error")
	}
}

func assertEquals(expected, actual interface{}) bool {
	if expected != actual {
		return false
	}
	return true
}
