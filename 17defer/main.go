package main

import "fmt"

func main() {
	defer fmt.Println("exec 3")
	fmt.Println("Defer")
	defer fmt.Println("exec 1")
	fmt.Println("exec 2")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
