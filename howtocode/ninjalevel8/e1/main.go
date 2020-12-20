package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Name string
		Age  int
	}
	u1 := User{
		"User1",
		25,
	}
	// u2 := user{
	// 	"user2",
	// 	26,
	// }
	// u3 := user{
	// 	"User3",
	// 	27,
	// }

	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
