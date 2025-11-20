package main

import "fmt"

func main() {
	fmt.Println("If else")

	cnt := 5
	var result string
	if cnt > 3 {
		result = "Regular"
	} else if cnt < 6 {
		result = "midregular"
	} else {
		result = "Irreguar"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Not less")
	}

	// if err!=nil{

	// }
}
