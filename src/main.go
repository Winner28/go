package main

import "algos"

//Saiyan struct
type Saiyan struct {
	Name  string
	Power int
}

func construct(name string, power int) Saiyan {
	return Saiyan{name, power}
}

func main() {
	stack := algos.InitStack([]int{1, 2, 3, 4}, "Regular stack!")
	stack.Push(10)
	algos.PrintStack(stack)
}
