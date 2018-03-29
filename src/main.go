package main

import (
	"algos"
	"fmt"
	"time"
)

func workWithStack() {
	fmt.Println("Work with Stack!")
	stack := algos.InitStack([]int{1, 2, 3, 4}, "Regular stack!")
	stack.Push(10)
	algos.PrintStack(stack)

}

func count(routine string) {
	for i := 0; i < 4; i++ {
		fmt.Println(i, routine)
		time.Sleep(time.Second * 1)
	}
}

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num)
	}
}

func main() {
	fmt.Println("...Entry point to main function...")
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5}
	go printCount(c)
	for _, value := range a {
		c <- value
	}
	time.Sleep(time.Second * 33)

}
