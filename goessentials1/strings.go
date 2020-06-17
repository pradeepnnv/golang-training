//This program gives overview of Strings in go language

package main

import "fmt"

func main() {
	var str string
	str = "This is a sample string containing a unicode "
	fmt.Println(str)
	fmt.Println("Length of the string str is ", len(str))
	fmt.Printf("This is a string whose type is %T with value %v\n", str, str)
	fmt.Printf("One of the characters at index 15 in the string str is of type %T with value %v\n", str[5], str[5])
	//Strings in go are immutable

	str = "0123456789"
	//Slice of a string
	//Starts from zero.
	fmt.Println(str[2:4])

}
