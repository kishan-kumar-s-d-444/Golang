package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "lco", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "lco", "bdd123", []string{"full-dev", "js"}},
		{"Angular Bootcamp", 299, "lco", "kis123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "lco",
		"tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("It is Valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Not valid")
	}

	//Cases to add data into key value
	var mydata map[string]interface{}
	json.Unmarshal(jsonData, &mydata)
	fmt.Printf("%#v\n", mydata)

	for k, v := range mydata {
		fmt.Printf("%v : %v\n", k, v)
		fmt.Printf("Type : %T\n", v)
	}
}
