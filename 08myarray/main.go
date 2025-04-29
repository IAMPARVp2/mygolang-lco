package main

import "fmt"

func main() {
	fmt.Println("welcome to the arrays using golang lecture")
	var fruitList [6]string

	fruitList[0] = "Apple"
	fruitList[1] = "orange"
	fruitList[2] = "banana"
	fruitList[4] = "grapes"

	fmt.Println("fruit list is: ", fruitList)
	fmt.Println("fruit list is: ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is: ", len(vegList))
}
