package main

import "fmt"

func main() {
	fmt.Println("Methods demo in golang")
	suj := User{"Sujeet", 32, "sujeet.go.dev", true}
	fmt.Println(suj)
	fmt.Printf("Sujeet details are%+v\n", suj)
	fmt.Printf("Name is %v and email is %v\n", suj.Name, suj.Email)

	suj.GetStatus()
	suj.NewEmail()
	fmt.Printf("Name is %v and email is %v\n", suj.Name, suj.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test.go.dev"
	fmt.Println("New email is", u.Email)
}
