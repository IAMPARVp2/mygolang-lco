package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com;3000/learn?coursename=djdfd&paymentid=rgegerger"

//above is the basic url strucutre

func main() {
	fmt.Println("welcome to handling urls in golang")
	fmt.Println(myurl)

	//parsing in this context mens that u have to extrat some information

	result, _ := url.Parse(myurl)

	fmt.Println(result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Println("the type of queryparams are: %T\n", qparam)

	fmt.Println("coursename")

	for _, val := range qparam {
		fmt.Println("param is : ", val)
	}
	//constructing a url
	partsOfUrl := &url.URL{ //alwys pss the reference of url here

		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tuusw",
		RawPath: "user=parbv",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)
}
