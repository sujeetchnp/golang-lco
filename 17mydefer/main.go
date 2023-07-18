package main

import "fmt"

func main() {
	fmt.Println("Defer demo in golang")

	defer fmt.Println("Two")
	defer fmt.Println("One")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

// Hello, myDefer(), World, One, Two

// myDefer = 0,1,2,3,4(stack) -- LIFO  = 43210

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
