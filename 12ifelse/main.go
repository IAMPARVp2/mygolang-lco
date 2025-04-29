package main

import "fmt"

func main() {
	fmt.Println("welcome to the if else")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 login counts"
	}
	fmt.Println(result)

	//==========================================
	if 9%2 == 0 {
		fmt.Println("it's even")
	} else {
		fmt.Println("it's odd")
	}
	//=================================================
	//another syntax for initializing and checking at same time
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("not less than 10")
	}

}
