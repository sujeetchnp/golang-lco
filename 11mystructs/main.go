package main

import "fmt"

func main() {
	fmt.Println("Structs demo in golang")
	suj := User{"Sujeet", 32, "sujeet.go.dev", true}
	fmt.Println(suj)
	fmt.Printf("Sujeet details are%+v\n", suj)
	fmt.Printf("Name is %v and email is %v\n", suj.Name, suj.Email)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
