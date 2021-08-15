package main

import "fmt"

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   24,
	}

	u2 := user{
		First: "Munroe",
		Age:   22,
	}

	u3 := user{
		First: "Sravs",
		Age:   29,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)
}
