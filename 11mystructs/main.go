package main

import "fmt"

func main() {
	fmt.Println("welcome to the structs")
	//no inheritance in go lang
	//no super or parent or child

	hitesh := User{"Hitesh", "jainv1071@gmail.com", true, 10}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)

	fmt.Printf("Name is %v  and email is %v.", hitesh.Email, hitesh.Name)
}

type User struct { //the first letter are capital to make them public so that we can use it at other places also
	Name   string
	Email  string
	Status bool
	Age    int
}
