package main

import (
	"fmt"
	"net/http"
)

func main() {
	contentType("https://www.linked.com")
	contentType("https://api.palestra.cccis.com")
}

func contentType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	cType := resp.Header.Get("Content-Type")
	if cType == "" {
		fmt.Println("No Content Type header in response")
	}
	fmt.Printf(`Content Type of the URL "%s" is "%s"
`, url, cType)
}
