package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getContentType(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "contenttypeheader: %s", err)
	}
	// fmt.Println(resp.Header.Get("Content-Type"))
	// return resp.Header.Get("Content-Type")
	ch <- resp.Header.Get("Content-Type")
}

func main() {
	start := time.Now()
	urls := []string{
		"https://www.google.com",
		"https://api.github.com",
		"https://api.twitter.com",
		"https://developer.twitter.com",
		"https://www.linkedin.com",
		"https://www.elastic.co",
		"https://prometheus.io",
		"https://home.pearsonvue.com/",
	}
	ch := make(chan string)
	for _, url := range urls {
		go getContentType(url, ch)
		fmt.Println(<-ch)
	}
	fmt.Println("Time taken: ", time.Now().Sub(start))
}
