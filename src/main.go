package main

import (
	"algos"
	"cli"
	"fmt"
	"interfaces"
)

type Fb interface {
	Do()
	private()
}

type str struct {
}

func (a str) Do() {

}

func (a str) private() {

}

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

func algosListRunner() {
	algos.RunList()
}

func main() {
	algosListRunner()
}
