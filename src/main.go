package main

import (
	"algos"
	"cli"
	"fmt"
)

func stackFunction() {
	fmt.Println("Work with Stack!")
	stack := algos.InitStack([]int{1, 2, 3, 4}, "Regular stack!")
	stack.Push(10)
	algos.PrintStack(stack)

}

func cliParser() {
	cli.ParseApp()
}

func main() {
	cliParser()
}
