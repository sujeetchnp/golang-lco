package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("It is an even number")
	} else {
		fmt.Println("It is an odd number")
	}
}
