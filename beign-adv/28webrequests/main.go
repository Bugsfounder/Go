package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - BF")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Length is: ", res.ContentLength)

	// better way to convert byte data into string or others, `string(content)` is not recommended for pro codes
	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"cousename":"Let's go with golang",
			"price":0,
			"platform":"youtube.com"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Bugs")
	data.Add("lastname", "Founder")
	data.Add("email", "bugsfounder2021@gmail.com")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)

	var responseString strings.Builder

	responseString.Write(content)
	// byteCount, _ := responseString.Write(content)

	fmt.Println(responseString.String())

}
