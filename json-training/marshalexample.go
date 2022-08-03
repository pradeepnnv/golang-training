package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	fmt.Println("hello")
	author := Author{
		"This is who wrote the book",
		36,
		true,
	}
	b := Book{
		Title:         "This is the book name",
		Author:        author,
		YearPublished: 2022,
	}
	fmt.Printf("%v\n", b)

	byteArray, err := json.MarshalIndent(b, "", "   ")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(string(byteArray))
}
