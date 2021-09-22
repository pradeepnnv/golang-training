package main

import (
	"log"
	"net/http"

	"dependencyinjection"
)

func main() {
	// dependencyinjection.Greet(os.Stdout, "John")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	dependencyinjection.Greet(w, "World")
}
