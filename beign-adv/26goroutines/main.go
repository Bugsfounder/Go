package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // usually variable is pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
		"https://bugsfounder.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
