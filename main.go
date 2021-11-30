package main

import (
	"log"
	"net/http"

	replace "dependencyinjection" ==> ./dependencyinjection /// Source code in /Users/ch-pnumb1-mbp/go/src/github.com/pradeepnnv/dependencyinjection
)

func main() {
	// dependencyinjection.Greet(os.Stdout, "John")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	dependencyinjection.Greet(w, "Pen")
}
