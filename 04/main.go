package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error-handling-04: %s", err)
		return
	}
	defer f.Close()
	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}
