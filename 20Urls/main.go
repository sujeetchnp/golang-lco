package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=sdfdsfh45464565"

func main() {
	fmt.Println("Urls demo in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
