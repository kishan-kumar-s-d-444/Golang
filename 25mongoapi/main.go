package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kishan-kumar-s-d-444/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB")
	r := router.Router()
	fmt.Println("Server is starting ...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000....")
}
