package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	inc := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println(inc)
		}()
	}
	wg.Wait()
	fmt.Println("Final value of inc is :", inc)
}
