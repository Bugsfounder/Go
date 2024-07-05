package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) // 02 -> day, 01 -> month, 2006 -> year, Monday --> day, casesensitive it has to be Monday not monday

	// learn more : https://go.dev/play/p/d_92jfpw8Xq

	createDate := time.Date(2020, time.August, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
