package main

import "fmt"

func main() {
	fmt.Println("MAPS")
	
	langs:=make(map[string]string)
	langs["JS"]="javascript"
	langs["RB"]="Ruby"
	langs["Py"]="Python"
	fmt.Println("List of all langs : ",langs)
	fmt.Println("Js means -- ",langs["JS"])

	delete(langs,"RB")
	fmt.Println(langs)

	//loops
	for key,val:=range langs{
		fmt.Printf("For key %v ,value is %v\n",key,val)
	}
}
