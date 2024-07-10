package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")

	// greeter // reference
	greeter() // call

	// Note: cannot create function inside function

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 5, 4, 5, 4, 3)
	fmt.Println("Pro result is: ", proResult)

	mulR, myMessage := proAdderMul(1, 4, 3, 23, 5)

	fmt.Println(mulR, myMessage)

}

// type of parameters and return types are compulsory
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// get any number of arguments or params (variadic functions)
func proAdder(values ...int) int {
	total := 0
	for _, values := range values {
		total += values
	}

	return total
}

// return multiple values
func proAdderMul(values ...int) (int, string) {
	total := 0
	for _, values := range values {
		total += values
	}

	return total, "Hello Pro Adder"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
