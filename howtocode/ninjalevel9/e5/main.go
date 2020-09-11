package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var inc int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			r := atomic.LoadInt64(&inc)
			fmt.Println(r)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of inc is :", inc)
}
