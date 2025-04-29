package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the slices lecture")

	var fruitList = []string{"apple", "banana", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mongo", "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 5)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 1465
	highScores[3] = 865
	highScores[4] = 945

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	//-----------------------------------

	var courses = []string{"reactjs", "javascript", "cpp", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
