package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u1 := user{
		"James",
		29,
	}
	u2 := user{
		"Sean",
		30,
	}
	u3 := user{
		"Bose",
		36,
	}
	users := []user{u1, u2, u3}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

type user struct {
	First string
	Age   int
}
