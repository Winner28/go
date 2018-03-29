package tests

import (
	"algos"
	"testing"
)

//TestQuickSort is working well
func TestQuickSort(t *testing.T) {
	actual := []int{5, 6, 10, 12, 3, 4}
	expected := []int{3, 4, 5, 6, 10, 12}
	quickSort(actual)
	if len(actual) != len(expected) {
		t.Error("Length not equals")
	}
	for index, value := range actual {
		if value != expected[index] {
			t.Error("Quick sort error. Not working!")
		}
	}
}

func quickSort(array []int) {
	algos.QuickSort(array, 0, len(array)-1)
}
