package main

import (
	"algos"
	"cli"
	"fmt"
	"interfaces"
)

func stack() {
	fmt.Println("Work with Stack!")
	stack := algos.InitStack([]int{1, 2}, "Regular stack!")
	stack.Push(10)
	_, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	algos.PrintStack(stack)

}

func cliPackage() {
	cli.ParseApp()
}

func interfacesPackage() {
	interfaces.Run()
}

func main() {
	//interfacesPackage()
	str := "hello!"
	t := str[1:2]
	t = "k"
	fmt.Println(t, str)
}
