package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println("HTTP_Requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type :%T\n", res)
	defer res.Body.Close()

	databytes,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content:=string(databytes)
	fmt.Println("Respnse:\n",content)
}
