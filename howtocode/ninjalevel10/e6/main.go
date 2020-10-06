package main

import "fmt"

func main() {
	c := make(chan int)

	//Push 100 integers into the channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	/*for v := range c {
		fmt.Println(v)
	}
	*/
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
