package main

import "fmt"

//number := 30000 //not allowed outside the method

//if u want to use it use it like this......

var numb = 3001 //or
var numb2 int = 3500

// ..........................................

const LoginToken string = "jain123" // this is a Public variable
/*as the first letter of variable name--"LoginToken"
signifies that this is Public,
so u can use it inside function or something*/

func main() {

	var username string = "parv"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45554584648561888
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable float64
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	//implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style

	numberOfUser := 30000.0 //this type of syntax is only allowed inside the method
	fmt.Println(numberOfUser)

	fmt.Println(numb)
	fmt.Println(numb2)

	fmt.Println(LoginToken) // using the public declared variable here from out side the function
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
