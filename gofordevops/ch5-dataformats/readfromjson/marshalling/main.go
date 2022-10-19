package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Record struct {
	Name string `json:"user_name`
	User string `json:"user"`
	ID   int
	Age  int `json:"-"`
}

func main() {
	rec := Record{
		Name: "Pradeep",
		User: "pradeepnnv",
		ID:   23,
	}
	// b, err := json.Marshal(rec)
	b, err := json.MarshalIndent(rec, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
