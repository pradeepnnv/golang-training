package main

import "fmt"

type person []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	fmt.Println("hello")
}
