package main

import "fmt"

func main() {
	x := `
	This is a raw literal string.
	So this can span multiple lines.
	This can even have ßß∂∑ˆ˙∆©¬¬˜œ∑ø¬¬¬åß∑∑™£££££££££™¡`
	fmt.Println(x)
}
