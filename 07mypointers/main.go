package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println("The value of ptr is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Address of memory is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Address of the memory is ", ptr)
	fmt.Println("New value is: ", *ptr)

}
