package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome")
	var num1 int = 2
	var num2 float64 = 4
	fmt.Println("The sum is: ", num1+int(num2))

	//random number two ways
	//1 mathematics
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//2 cryptography
	randnum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randnum)
}
