package main
import "fmt"

const Logintoken string="12user0kishan"


func main(){
	var username string="kishan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n",username)

	var smallval uint8=255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T\n",smallval)

	var smallfloat float64=255.45454545454449283402394289
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T\n",smallfloat)

	//No-value
	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("Variable of type :  %T \n",anothervar)

	//implicit type
	var website="hi.com"
	fmt.Println(website)

	//No var keyword
	noofuser :=4
	fmt.Println(noofuser)

	fmt.Println(Logintoken)
	fmt.Printf("Variable of type :  %T \n",Logintoken)
}
