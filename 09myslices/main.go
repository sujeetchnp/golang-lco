package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the tutorial of Slices in golang")

	vegList := []string{"tomato", "potato", "onion"}
	vegList = append(vegList, "eggplant", "spinach")
	fmt.Println("Appended vegList slice", vegList)

	vegList = append(vegList[1:3])
	fmt.Println(vegList)

	highScores := make([]int, 4)

	highScores[0] = 445
	highScores[1] = 554
	highScores[2] = 565
	highScores[3] = 900
	// highScores[4] = 800

	fmt.Println(highScores)

	highScores = append(highScores, 800, 1000)

	fmt.Println("After append", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted highScores", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	courses := []string{"golang", "java", "reactjs", "angularjs", "c++", "swift"}
	fmt.Println("Courses", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After splitting index", courses)

}
