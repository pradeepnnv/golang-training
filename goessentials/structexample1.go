// This is an example for structs in Golang

package main

import "fmt"

// Trade is trade in stocks
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 2, 106.4, false}
	fmt.Println(t1)
	t2 := Trade{
		Symbol: "GOOG",
		Volume: 5,
		Price:  205.8,
		Buy:    true,
	}
	fmt.Println(t2.Buy)
	t3 := Trade{}
	fmt.Println(t3)

	fmt.Printf("%+v\n", t3)
}
