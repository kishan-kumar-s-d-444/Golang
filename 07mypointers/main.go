package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// var ptr *int
	// fmt.Println("Value of ptr is : ",ptr)

	num:=23
	var ptr= &num
	fmt.Println("Value of actual pointer is : ",ptr)
	fmt.Println("Value of actual pointer is : ",*ptr)

	*ptr=*ptr*4
	fmt.Println("New val is: ",num)

}
