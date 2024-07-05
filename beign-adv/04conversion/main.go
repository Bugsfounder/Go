package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to our pizza app"
	fmt.Println(welcome)
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// this will also work but not preferable
	// inp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Println(inp)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	// conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
 