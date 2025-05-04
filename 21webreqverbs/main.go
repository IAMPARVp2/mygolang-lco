package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to the web verb video- LCO")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get" //This is the URL where your local HTTP server should be running and listening on /get route.

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	//This line closes the response body automatically when the function exits
	// (a good practice to avoid memory leaks or open connections).
	fmt.Println("Status code: ", response.StatusCode)          //gives you the HTTP status (e.g., 200 OK, 404 Not Found).
	fmt.Println("Content length is: ", response.ContentLength) //tells you the number of bytes in the response body.

	//ONE WAY-
	content, _ := io.ReadAll(response.Body)
	fmt.Println(content)
	fmt.Println(string(content))

	//OTHER WAY-
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	
		{
			"coursename": "Lets og with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "jainvever@fer")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
