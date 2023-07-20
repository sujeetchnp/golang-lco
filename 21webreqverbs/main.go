package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"
	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bytecount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("ByteCount is: ", bytecount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myURL = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"price":450,
			"paltform":"udemy.com"
		}
	`)
	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myURL = "http://localhost:8000/post"

	data := url.Values{}
	data.Add("firstname", "sujeet")
	data.Add("lastname", "chaudhary")
	data.Add("email", "sujeet@go.dev")

	response, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
