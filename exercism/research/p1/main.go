package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := 60
	var answer string

	if x%144 == 0 {
		// answer += strconv.Itoa(x/144) + " gross or "
		answer = strings.Join([]string{answer, strconv.Itoa(x / 144), "gross or"}, " ")
	}
	if x%20 == 0 {
		// answer += strconv.Itoa(x/20) + " score or "
		answer = strings.Join([]string{answer, strconv.Itoa(x / 20), "score or"}, " ")
	}
	if x%12 == 0 {
		// answer += strconv.Itoa(x/12) + " dozen or "
		answer = strings.Join([]string{answer, strconv.Itoa(x / 12), "dozen or"}, " ")
	}
	// answer += strconv.Itoa(x)
	answer = strings.Join([]string{answer, strconv.Itoa(x)}, " ")
	fmt.Println(answer)
}
