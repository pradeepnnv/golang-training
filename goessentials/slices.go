package main

import (
	"Fmt"
)

func main() {
	loons := []string{"flat", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)
	fmt.Println("Length of this slice is ...:", len(loons))
	fmt.Println(loons[:3])
	fmt.Println(loons[2:])

	//String loops
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}
	//Single value range
	for i := range loons {
		fmt.Println(i)
	}
	//Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}
	//Double valued range, ignoring index
	for _, name := range loons {
		fmt.Println(name)
	}
}
