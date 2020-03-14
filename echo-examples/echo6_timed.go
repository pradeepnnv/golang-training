//Echo prints command line arguments and measures the time taken
//Excercise 1.3

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var elapsed time.Duration
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	elapsed = time.Since(start)
	fmt.Println("Echo took %s", elapsed)
}
