package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to file in golang")
	content := "my name is parv jain"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("text data inside the file is \n", string(databyte))

}
