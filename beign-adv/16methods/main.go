package main

import (
	"fmt"
)

func main() {

	bugs := User{"Bugs Founder", "bugsfounder2021@gmail.com", true, 20}
	fmt.Println(bugs)

	bugs.GetStatus()
	bugs.NewMail()
	fmt.Println(bugs.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // private 'first letter small'
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
