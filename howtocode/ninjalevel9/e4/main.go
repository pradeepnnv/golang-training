package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	inc := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			v := inc
			// runtime.Gosched()
			v++
			inc = v
			m.Unlock()
			fmt.Println(inc)
		}()
	}
	wg.Wait()
	fmt.Println("Final value of inc is :", inc)
}
