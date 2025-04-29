package main

import "fmt"

func main() {
	fmt.Println("welcome to defer")
	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("two")
	mydefer()

}

//output
/*
welcome
hello
two
43210world
*/

func mydefer() {
	for i := 0; i < 5; i++ {

		defer fmt.Println(i)
	}
}
