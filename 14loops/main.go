package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("For index %v, the day is %v\n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			// break         // terminates the loop completly when rougueValue becomes 5
			rougueValue++
			continue // terminates the the only loop when rougueValue becomes 5 and other will continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++

	}
lco:
	fmt.Println("Jumping at LearnCodeonline.in")

}
