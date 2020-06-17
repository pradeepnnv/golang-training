//Program to identify all even ended numbers when two four digit numbers are multiplied

package main

import "fmt"

func main() {
	//Count of even ended numbers
	count := 0
	var numstr string
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			numstr = fmt.Sprintf("%v", i*j)
			if numstr[0] == numstr[len(numstr)-1] {
				count++
			}
		}
	}

	fmt.Printf("%q", fmt.Sprintf("%v", count))

}
