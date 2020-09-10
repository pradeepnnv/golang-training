package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Users type represents a slice of users
type Users []user

// Len returns the length of the slice u
func (u Users) Len() int {
	return len(u)
}

// Swap swaps the users i,j
func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Less determines if u[i] or u[j] is less
func (u Users) Less(i, j int) bool {

	return u[i].Age < u[j].Age
}
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(Users(users))

	for _, u := range users {
		fmt.Println("First name:", u.First)
		fmt.Println("Last name:", u.Last)
		fmt.Println("Age:", u.Age)
		fmt.Println("Sayings:")
		sort.Strings(u.Sayings)
		for _, saying := range u.Sayings {
			fmt.Println("\t\t\t", saying)
		}
	}
}
