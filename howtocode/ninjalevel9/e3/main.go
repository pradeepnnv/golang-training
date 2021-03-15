package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	i := 0
	wg.Add(5)
	for x := 0; x < 5; x++ {
		go func() {
			i++
			fmt.Println("Value of i in iteration " , x  i)
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(i)
}
