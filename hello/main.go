package main

import (
	"fmt"
	"log"

	"go.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// slice of names.
	names := []string{"Bugs", "Samantha", "Kosu"}

	// request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal("err")

	}

	// if no error was returned, print the returned message to the console
	fmt.Println(messages)
}
