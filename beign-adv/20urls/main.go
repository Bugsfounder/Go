package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.bugsfounder.com/learn?coursename=go&payment_id=asdfa4345-bugs"

func main() {
	fmt.Println("Welcome to handling urls")

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawPath: ", result.RawPath)
	fmt.Println("RawQuery: ", result.RawQuery)
	fmt.Println("User: ", result.User)

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["payment_id"])

	for key, value := range qparams {
		fmt.Println(key, value)
	}

	// url must be always a reference '&url' not 'url'
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "bugsfounder.com",
		Path:    "/go",
		RawPath: "user=bugs",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
