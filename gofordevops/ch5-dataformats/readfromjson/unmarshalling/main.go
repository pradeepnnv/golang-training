package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(b, &data); err != nil {
		log.Fatal(err)
	}

	v, ok := data["user"]
	if !ok {
		log.Fatal("json does not contain key 'user'")
	}

	switch user := v.(type) {
	case string:
		fmt.Println("User is ", user)
	}
}
