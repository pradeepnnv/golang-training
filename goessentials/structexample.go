package main

import "fmt"

//Trade is a struct
type Trade struct {
	Symbol string  // Stock Symbol
	Volume int     // Number of Shares
	Price  float64 // Value of the Stock
	Buy    bool    // true if buy Trade
}

func main() {

	t1 := Trade{"MSFT", 5, 23.46, false}
	fmt.Println(t1)
  fmt.Printf("%+v\n",t1)

  t2 := Trade{
    Symbol: "MSFT",
    Volume: 5,
    Price: 1232.4,
    Buy: true,
  }

fmt.Printf("%+v\n",t2)
t3:= Trade{}
fmt.Printf("%+v\n",t3)

}
