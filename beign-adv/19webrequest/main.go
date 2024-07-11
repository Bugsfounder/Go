package main

import (
	"fmt"
	"io"
	"net/http"
)

// dummy apis: https://jsonplaceholder.typicode.com/
const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to web request")

	res_content, err := request(url)

	checkNilError(err)

	fmt.Println(res_content)

}

func request(url string) (data string, err error) {
	res, err := http.Get(url)

	if res.StatusCode == 200 {

		databyte, err := io.ReadAll(res.Body)

		checkNilError(err)

		return string(databyte), nil
	}

	return "", err
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
