package main

import "fmt"

func main() {
	fmt.Println("Structs")

	kishan := User{"Kishan", "k@gmail.com", true, 21}
	fmt.Println(kishan)
	fmt.Printf("Details are : %+v", kishan)
	fmt.Printf("Email = %v", kishan.Email)
	kishan.GetStatus()
	kishan.NewMail()
	fmt.Printf("Email = %v", kishan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@email.com"
	fmt.Println("Email changed to", u.Email)
}
