package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com"

func main() {
	fmt.Println("URLS")
	fmt.Println(myurl)

	//parsing
	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("Type of Query Params are:%T\n", qparams)

	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	//construt url
	fmt.Println("Construct URL")
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=kishan",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
