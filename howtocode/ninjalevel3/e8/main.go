package main

import "fmt"

func main() {
	i := 8
	switch {
	case i == 4:
		{
			fmt.Println("Value of i is 4")
		}
	case i == 8:
		{
			fmt.Println("Value of i is 8")
		}
	default:
		{
			fmt.Println("Value of is is not matching")
		}
	}

}
