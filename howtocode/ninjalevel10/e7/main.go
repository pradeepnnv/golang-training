package main

import "fmt"

func main() {
	c := make(chan int)
	for i, v1 := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
	close(c)
	receive(c)
}

func receive(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
