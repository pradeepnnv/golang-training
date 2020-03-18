//This program calculates mean of two numbers

package main

import (
"fmt"
)

func main() {
var x float64
var y float64
x = 25.934343
y = 48.2323234399999
fmt.Printf("x=%v , type of %T\n",x,x)
fmt.Printf("y=%v, type of %T\n",y,y)

mean := (x+y)/2
fmt.Printf("mean=%v, type of %T\n",mean,mean)

}
