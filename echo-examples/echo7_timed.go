//Prints all command line arguments along with duration
//Excercise 1.3

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var elapsed time.Duration
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start)
	fmt.Println("Echo took %s", elapsed)
}
