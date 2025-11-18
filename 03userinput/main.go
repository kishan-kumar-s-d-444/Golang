package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Come On"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating :")

	//comma ok || err err syntax

	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for rating,",input)
	fmt.Printf("Type of rating is %T",input)
}
