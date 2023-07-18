package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch case in golang")

	//rand.Seed(time.Now().UnixNano())  // this is deprecated from 1.18, no need now

	diceNumber := rand.Intn(6) + 1 // use rand.Intn directly
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Rolled 1, so you can open")
	case 2:
		fmt.Println("Rolled 2, you can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("Rolled 3, you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("Rolled 4, you can move 4 spot")
	case 5:
		fmt.Println("Rolled 5, you can move 5 spot")
	case 6:
		fmt.Println("Rolled 6, you can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}

}
