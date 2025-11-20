package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greeter()
	res := adder(3, 5)
	fmt.Println("Result = ", res)
	prores, messages := proAdder(3, 5, 6, 7)
	fmt.Println("Result proadder", prores)
	fmt.Println("Message", messages)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "result function"
}

func greeter() {
	fmt.Println("Welcome")
}
