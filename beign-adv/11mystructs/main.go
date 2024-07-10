package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	// no inheritance in go lang; No super or parent
	bugs := User{"Bugs Founder", "bugsfounder2021@gmail.com", true, 20}
	fmt.Println(bugs)
	fmt.Printf("Bugs details are: %+v\n", bugs)
	fmt.Printf("Name is %v and email is %v\n", bugs.Name, bugs.Email)
	
}

type User struct {
	Name   string 
	Email  string
	Status bool
	Age    int
}
