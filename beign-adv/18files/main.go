package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - BugsFounder.com" // content
	filename := "./mybffile.txt"                              // filename

	writeFile(filename, content)

	readFile(filename) // reading file './mybffile.txt'

}

// readFile function to read a content of a file and prints the content of that file into the terminal
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) // deprecated
	// default data type of ReadFile func is Byte

	checkNilError(err)

	fmt.Println("Text Data inside the file is \n", string(databyte))
}

func writeFile(filename string, content string) {
	// Creating a file
	file, err := os.Create(filename)
	// this logic is repeating to many times so it is good to wrap it into a function.
	// if err != nil {
	// 	panic(err)
	// }
	// checking error
	checkNilError(err)

	// writing to the file
	length, err := io.WriteString(file, content)

	// checking error
	checkNilError(err)

	fmt.Println(length)

	// close file at the end of the program 'end of the main function'
	defer file.Close()

}

// function to check err so we don't need to repeat the same logic again and again.
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
