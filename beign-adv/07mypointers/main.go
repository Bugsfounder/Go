package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("value of pointer is:", ptr)

	// pointer -> *, reference -> &
	myNumber := 23
	// var ptr *int = &myNumber
	var ptr = &myNumber
	// var ptr *int = myNumber // error
	fmt.Println("Value of actual pointer is:", ptr)
	fmt.Println("Value of actual pointer is:", *ptr)

	// updating variable value through pointer
	*ptr = *ptr + 3
	fmt.Println("New Value is:", *ptr)

}
