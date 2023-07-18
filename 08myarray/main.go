package main

import "fmt"

func main() {
	fmt.Println("Welcome to the tutorial of arrays in golang")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"
	fruitList[4] = "Guava"

	fmt.Println("Fruits are: ", fruitList)
	fmt.Println("Length of fruitList array is ", len(fruitList))

	vegList := [4]string{"tomato", "potato", "onion"}
	fmt.Println("veggies are: ", vegList)
	fmt.Println("Length of vegList array is ", len(vegList))

}
