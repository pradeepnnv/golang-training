//Program which contains a function which returns contentType header in response to a URL.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	cType, err := contentType("https://www.google.com")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(cType)
	}
}

func contentType(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		return "", fmt.Errorf("Content Type header is not available")
	}
	return cType, nil
}
