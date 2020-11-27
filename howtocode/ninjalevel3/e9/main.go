package main

import "fmt"

func main() {
	favSport := "Mario"

	switch favSport {
	case "Football":
		fmt.Println("Favourite sport is Football")
	case "Tennis":
		fmt.Println("Favourite sport is Tennis")
	case "Mario":
		fmt.Println("Favourite sport is Mario")
	default:
		fmt.Println("Looks like he does not like any games")
	}
}
