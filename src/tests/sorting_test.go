package tests

import (
	"algos"
	"testing"
)

//TestQuickSort is working well
func TestQuickSort(t *testing.T) {
	array := []int{5, 6, 10, 12, 3, 4}
	sortArray := []int{3, 4, 5, 6, 10, 12}
	quickSort(array)
	if len(array) != len(sortArray) {
		t.Error("Length not equals")
	}
	for index, value := range array {
		if value != sortArray[index] {
			t.Error("Quick sort error. Not working!")
		}
	}
}

func quickSort(array []int) {
	algos.QuickSort(array, 0, len(array)-1)
}
