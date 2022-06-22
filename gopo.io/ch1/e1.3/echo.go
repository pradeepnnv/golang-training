package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	/*
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
				s += sep + os.Args[i]
				sep = " "
			}
			fmt.Println(s)
	*/

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
