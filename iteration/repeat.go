package iteration

import "strings"

// Repeat repeats the `input` string `count` number of times.
// using standard library
// most efficient
func Repeat(i string, c int) string {
	return strings.Repeat(i, c)
}

// comparitively efficient version.
// func Repeat(input string, count int) string {
// 	var sb strings.Builder
// 	for i := 0; i < count; i++ {
// 		sb.WriteString(input)
// 	}
// 	return sb.String()
// }

// very inefficient version
// func Repeat(input string, count int) string {
// 	var output string

// 	for i := 0; i < count; i++ {
// 		output += input
// 	}
// 	return output
// }
