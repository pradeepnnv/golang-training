//Program which contains a function which returns contentType header in response to a URL.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.cccis.com",
		"https://mapquest.com",
		"https://api.github.com",
	}
	for _, url := range urls {
		contentType(url)
	}
}

//Using goroutines
func contentType(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	defer resp.Body.Close()
	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		fmt.Println("Content Type header is not available")
	} else {
		fmt.Println(url, "===>", cType)
	}
}
