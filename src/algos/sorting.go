package algos

import "fmt"

//SelectionSort sort
func SelectionSort(arr []int) {
	for i, value := range arr {
		index := i
		for j := i; j < len(arr); j++ {
			if arr[j] > value {
				value = arr[j]
				index = j
			}
		}
		if index != i {
			swap(arr, i, index)
		}
	}
}

//InsertionSort sort
func InsertionSort(arr []int) {
	for i, value := range arr {
		j := i - 1
		for j >= 0 && arr[j] < value {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = value
	}
}

//QuickSort sort
func QuickSort(arr []int, left, right int) {
	l := left
	r := right
	pivot := l + (r-l)/2

	for l <= r {
		for arr[pivot] > arr[l] {
			l++
		}
		for arr[pivot] < arr[r] {
			r--
		}

		if l <= r {
			swap(arr, l, r)
			l++
			r--
		}

		if left < r {
			QuickSort(arr, left, r)
		}
		if right > l {
			QuickSort(arr, l, right)
		}

	}
}

//MergeSort sort
func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	middle := left + (right-left)/2
	MergeSort(arr, 0, middle)
	MergeSort(arr, middle+1, right)
	merge(arr, left, middle, right)
}

func merge(array []int, left, middle, right int) {
	l := array[0:middle]
	r := array[middle+1 : right]
	i, j, p := 0, 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			array[p] = l[i]
			i++
		} else {
			array[p] = r[j]
			j++
		}
		p++
	}

	//TODO: todo

	for i < len(l) {
		array[p] = l[i]
		p++
		i++
	}

	for j < len(r) {
		array[p] = r[j]
		j++
		p++
	}
}

func swap(arr []int, index1, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

//Print array
func Print(arr []int) {
	for i, v := range arr {
		fmt.Printf("Index: %d *** Value: %d \n", i, v)
	}

}
