package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// Basic / old
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// Get index in i and iterate
	for i := range days {
		// i is index
		fmt.Println(days[i])
	}

	// kind of for-each
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 1
	// kind of while loop
	for rougueValue < 10 {
		fmt.Println("Value is : ", rougueValue)
		// ++rougueValue // error
		rougueValue++
	}

	rougueValue = 1
	// break
	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			break
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}
	// continue
	rougueValue = 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

	// goto
	rougueValue = 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto bf
		}
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}
bf:
	fmt.Println("Jumping at bugsfounder.com")

}
