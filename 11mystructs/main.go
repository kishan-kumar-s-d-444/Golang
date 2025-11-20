package main

import "fmt"

func main() {
	fmt.Println("Structs")

	kishan:=User{"Kishan","k@gmail.com",true,21}
	fmt.Println(kishan)
	fmt.Printf("Details are : %+v",kishan)
	fmt.Printf("Name = %v",kishan.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
