package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	proLangs := make(map[string]string)

	proLangs["JS"] = "JavaScript"
	proLangs["J"] = "Java"
	proLangs["RJs"] = "React JavaScript"
	proLangs["RB"] = "Ruby"
	proLangs["Ajs"] = "Angular JavaScript"

	fmt.Println("List of programming languages ", proLangs)
	fmt.Println("For key value RB is", proLangs["RB"])

	delete(proLangs, "RB")
	fmt.Println("List of programming languages after deleting RB", proLangs)

	// loops
	for key, value := range proLangs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// for key, _ := range proLangs {
	// 	fmt.Println("keys are", key)
	// }

	for _, value := range proLangs {
		fmt.Println("values are", value)
	}
}
