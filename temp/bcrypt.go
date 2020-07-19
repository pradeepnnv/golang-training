package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Fprint(os.Stdout, "bcrypt: %s", err)
	}
	fmt.Println("Original pwd:", s)
	fmt.Printf("Encrypted pwd: %s", bs)
}
