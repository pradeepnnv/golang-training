//Program to fetch contents at a given URL
//URL is passed as a command line argument
//Excercise 1.7 in GOPL

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch v1: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch v1: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
