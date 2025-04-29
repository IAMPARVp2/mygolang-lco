package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "friday", "Saturday", "Monday"} //list or slice

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Print(days[i], " ")
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("index is %v and value is %v\n", index, days)
	}

	for _, days := range days {
		fmt.Printf("index is and value is %v\n", days)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			break
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("jumping to my website parv jain")
}
