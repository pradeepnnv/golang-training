//Program which contains a function which returns contentType header in response to a URL.

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.cccis.com",
		"https://mapquest.com",
		"https://api.github.com",
	}
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			contentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()

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
