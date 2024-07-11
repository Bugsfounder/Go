package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"_"`              // ignored, i don't want to reveal password
	Tags     []string `json:"tags,omitempty"` // if nil then leave it empty
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
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
