package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fSeq := []int{
		1,
		2,
	}
	sumEvenTerms := 2
	for {
		nextVal := fSeq[len(fSeq)-1] + fSeq[len(fSeq)-2]
		if nextVal >= 4000000 {
			break
		}
		fSeq = append(fSeq, nextVal)
		if nextVal%2 == 0 {
			sumEvenTerms += nextVal
		}
	}
	fmt.Println("Sum of Even Terms is :", sumEvenTerms)

	secs := time.Since(start).Seconds()
	fmt.Printf("Time taken: %.2fs seconds\n", secs)
}
