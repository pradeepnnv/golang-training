package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	var answer string

	if x%144 == 0 {
		answer += strconv.Itoa(x/144) + " gross or "
	}
	if x%20 == 0 {
		answer += strconv.Itoa(x/20) + " score or "
	}
	if x%12 == 0 {
		answer += strconv.Itoa(x/12) + " dozen or "
	}
	answer += strconv.Itoa(x)
	fmt.Println(x%144, x/144)
	fmt.Println(x%20, x/20)
	fmt.Println(x%12, x/12)
	fmt.Println(answer)
}
