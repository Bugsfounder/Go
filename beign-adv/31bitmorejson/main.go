package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // ignored, i don't want to reveal password
	Tags     []string `json:"tags,omitempty"` // if nil then leave it empty
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	bfcourses := []course{
		{"ReactJS Bootcamp", 266, "LearnCodeOnline.in", "asdeoi@3342", []string{"web", "js", "dev"}},
		{"MERN Bootcamp", 443, "LearnCodeOnline.in", "bets#433", []string{"web", "js", "dev", "mern"}},
		{"Angular Bootcamp", 456, "LearnCodeOnline.in", "w3erwsd23543", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(bfcourses)
	// finalJson, err := json.MarshalIndent(bfcourses, "BF", "\t")
	finalJson, err := json.MarshalIndent(bfcourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %s\n", finalJson)

}

func DecodeJson() {

	// this is just a dummy string json data we have converted it into byte so it looks like coming for web because web response are in byte.
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 443,
		"website": "LearnCodeOnline.in",
		"_": "bets#433",
		"tags": ["web","js","dev","mern"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value (map)
	var myOnlineData map[string]interface{} // we don't know what is the type of value is so we use interface

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n", key, value, value)
	}
}
