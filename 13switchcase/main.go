package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6) + 1
	fmt.Println("Dice rolled to : ", no)
	switch no {
	case 1:
		fmt.Println("it is 1")
	case 2:
		fmt.Println("It is 2")
	case 3:
		fmt.Println("it is 3")
		fallthrough
	case 4:
		fmt.Println("it is 4")
		fallthrough
	case 5:
		fmt.Println("it is 5")
	case 6:
		fmt.Println("it is 6 roll again")
	default:
		fmt.Println("You are wrong")
	}
}
