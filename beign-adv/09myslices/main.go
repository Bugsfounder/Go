package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	fruitList = fruitList[1:3]
	fruitList = fruitList[0:3]
	fruitList = fruitList[0:]
	fruitList = fruitList[0:5]
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 944
	highScores[2] = 454
	highScores[3] = 876
	//	highScores[4] = 777 // error

	highScores = append(highScores, 555, 666, 321) // ok
	fmt.Println(highScores)
	fmt.Println(len(highScores))

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

}
