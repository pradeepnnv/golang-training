package main

//Echo1 prints it's command line arguments.

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)

	fmt.Printf("Time taken is %v\n", elapsed)
}
