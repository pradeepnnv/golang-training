package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// User contains count of public repos of a user and his user login
type User struct {
	PublicRepos int    `json:"public_repos"`
	Name        string `json:"login"`
}

func main() {
	url := "https://api.github.com/users/pradeepnnv"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "githubapi_challenge: %s", err)
	}
	defer resp.Body.Close()
	// fmt.Println(resp)
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		fmt.Fprintf(os.Stderr, "githubapi_challenge: %s", err)
	}
	fmt.Println(*user)
}
