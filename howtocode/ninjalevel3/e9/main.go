package main

import "fmt"

func main() {
	favSport := "foosball"
	switch {
	case "football" == favSport:
		fmt.Println("You are wrong")
	case "foosball" == favSport:
		fmt.Println("You are right")
	default:
		fmt.Println("Whaaat!!!")
	}
}
