//maps explains maps concepts in golang
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1807.84,
		"AAPL": 252.86,
		"GOOG": 1118.06,
		"MSFT": 146.57,
	}
	fmt.Println("Length of stocks map is :", len(stocks))

	//Get a value
	fmt.Println(stocks["AMZN"])
	//Get a nonexistent value
	fmt.Println(stocks["TSLA"])
	//Check for nonexistent value
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA value not found")
	} else {
		fmt.Println(value)
	}
	//Set example
	stocks["TSLA"] = 432.0
	fmt.Println(stocks)

	//Delete example
	delete(stocks, "AAPL")
	fmt.Println(stocks)

	//Single value for is on keys
	for key := range stocks {
		fmt.Println(key)
	}
	//Multi value is key,value
	for key, value := range stocks {
		fmt.Printf("Value of %s is %f\n", key, value)
	}
}
