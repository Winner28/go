package main

import (
	"algos"
	"cli"
	"fmt"
	"interfaces"
	"multithreading"
	"multithreading/storage"
	"web/server"
	"web/simple"
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

func algosListRunner() {
	algos.RunList()
}

func startSimpleServer() {
	simple.StartSimpleServer(8181)
}

func goodStorage() {
	storage.RunStorageProg()
}

func runSimpleEcho() {
	multithreading.RunSimpleEcho()
}

func main() {
	server.StartServer()
}
