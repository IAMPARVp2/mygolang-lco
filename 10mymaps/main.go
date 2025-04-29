package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")

	languages := make(map[string]string)
	languages["JS"] = "javaScript"
	languages["RB"] = "Ruby"
	languages["CPP"] = "cPlusPlus"
	languages["PY"] = "python"

	fmt.Println("list of all languages:", languages)
	fmt.Println("JS full form is", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages:", languages)

	//loops in go lang

	//for loop
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
		//%v is for value
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
		//%v is for value
	}
}
