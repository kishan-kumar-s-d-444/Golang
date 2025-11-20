package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Sataurday"}
	
	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	for i:=range days{
		fmt.Println(days[i])
	}

	fmt.Println("-----------")

	for index,day :=range days{
		fmt.Printf("index : %v and value %v\n",index,day)
	}

	fmt.Println("-----------")
	rval:=1
	for rval<10{
		if rval==2 {
			goto hi
		}
		if rval==7 {
			break
		}
		if rval==3 {
			rval++
			continue
		}
		fmt.Println(rval)
		rval++
	}
	hi:
		fmt.Println("Jumping statement")
	fmt.Println("-----------")

}
