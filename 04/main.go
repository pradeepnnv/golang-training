package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*	This program creates a file `names.txt` which contains the two names.
 */
func main() {
	f, err := os.Create("names.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error-handling-04: %s", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond\n")
	io.Copy(f, r)

	r = strings.NewReader("Swami Vivekananda")
	io.Copy(f, r)
}
