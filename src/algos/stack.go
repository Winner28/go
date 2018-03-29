package algos

import "fmt"

//Stack structure
type Stack struct {
	array []int
	name  string
}

//InitStack fabrics method
func InitStack(array []int, name string) Stack {
	return Stack{
		array: array,
		name:  name,
	}

}

//Push to stack
func (s *Stack) Push(value int) Stack {
	//actually we dont need this all, I wrote this just for fun =)
	if len(s.array) == cap(s.array) {
		fmt.Println("The capacity has come to an end. Expand an array! (cap before: ", cap(s.array), " )")
		arr := make([]int, len(s.array), cap(s.array)*2)
		copy(arr, s.array)
		arr = append(arr, value)
		s.array = arr
		fmt.Println("Capacity after: ", cap(s.array))
	}
	return *s
}

//Pop from stack
func (s *Stack) Pop() (int, error) {

	return 100, nil
}

//PrintStack prints stack
func PrintStack(s Stack) {
	fmt.Println("Stack name: ", s.name)
	for index, value := range s.array {
		fmt.Printf("Index: %d *** Value: %d \n", index, value)
	}
}
