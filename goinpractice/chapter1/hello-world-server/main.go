package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `నమస్కారము`)
}
