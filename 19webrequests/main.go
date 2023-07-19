package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("How Web request works in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibilty to close the connection

	datatypes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(datatypes)
	fmt.Println(content)

}
