package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TIME PACKAGES")

	presentTime := time.Now()
	fmt.Println("Current time: ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(1004, time.August, 2, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
