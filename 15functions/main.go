package main

import "fmt"

func main() {
	fmt.Println("welcome to the functions")
	greeter()
	greeter2()

	result := adder(3, 5)
	fmt.Println("the result after adding the two numbers is:", result)

	proRes, myMessage := proAdder(2, 5, 6, 9, 4)
	fmt.Println("Pro result is:", proRes)
	fmt.Println("My pro message is:", myMessage)
}

func greeter() {
	fmt.Println("Parv from go lang")
}

func greeter2() {
	fmt.Println("another method")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) { //takes values , ...int means-> take all the values as integers
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hi pro result function"
}
