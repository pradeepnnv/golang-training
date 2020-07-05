//Program to retreive content at a URL
//Excercise 1.8 GOPL

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			// fmt.Println("URL does not have http:// prefix. Aborting...")
			// os.Exit(2)
			url = "http://" + url
			fmt.Println("Provided url does not have any protocol. Defaulting to http and using ", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch v2: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch v2: error getting %s: %v\n", url, err)
		}
		fmt.Printf("%s\n", b)
	}
}
