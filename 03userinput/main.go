package main

import (
	"bufio" //For buffered input/output — helps you read text from the user.
	"fmt"
	"os" //For interacting with the operating system (like reading input).
)

func main() {
	welcome := "Welcome to user input" //You create a string variable called welcome.
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //NewReader is a function in the bufio package.
	/*You create a reader that reads input
	from the keyboard (standard input) and
	bufio.NewReader(os.Stdin) means:
	“Get ready to read what the user types.”*/
	fmt.Println("Enter the rating of the pizza: ")

	//comma ok // err err

	input, _ := reader.ReadString('\n') //This line actually reads what the user types.
	//ReadString('\n') = Read until the user hits Enter (\n = newline).
	//The _ means you're ignoring any error that might happen (for now).
	//input stores the text the user typed (including the newline \n).
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of this rating is %T", input)
}
