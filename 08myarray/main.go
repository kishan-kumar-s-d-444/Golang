package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruits [4] string
	fruits[0]="Apple"
	fruits[1]="Tomato"
	fruits[3]="Peach"
	
	fmt.Println("Fruits are",fruits)
	fmt.Println("Fruits length are",len(fruits))

	var veggies =[3]string{"p","b","m"}

	fmt.Println("Veggies are : ",veggies)
	fmt.Println("Veggies length : ",len(veggies))

}
