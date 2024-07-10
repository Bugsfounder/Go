package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in go lang")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all language: ", language)
	fmt.Println("Js Shorts for: ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all language: ", language)

	// loops are interesting in golang
	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// comma ok
	for _, value := range language {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
