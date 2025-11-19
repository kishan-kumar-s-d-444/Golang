package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruits = []string{"A", "B", "C"}
	fmt.Printf("Type of fruits : %T\n", fruits)

	fruits = append(fruits, "D", "E")
	fmt.Println("FFruits : ", fruits)

	fruits = append(fruits[1:])
	fmt.Println("Fruits", fruits)

	scores := make([]int, 4)
	scores[0] = 234
	scores[1] = 945
	scores[2] = 465
	scores[3] = 867
	fmt.Println("Scores", scores)
	// scores[4]=777
	// fmt.Println("Scores",scores)
	scores = append(scores, 555, 666, 321)
	fmt.Println("Scores : ", scores)
	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)

	//Remove value from slices based on index

	var courses = []string{"js", "py", "cpp", "go"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
