package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()
	fmt.Println("Adder function result", adder(2, 5))
	variadicFunc, myMessage := varAdder(4, 5, 6, 7, 8, 9)
	fmt.Println("Variadic result:", variadicFunc, ";", myMessage)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func greeterTwo() {
	fmt.Println("Another function")
}

func adder(num1, num2 int) int {
	return num1 + num2
}

func varAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Varidic message: Success"
}
