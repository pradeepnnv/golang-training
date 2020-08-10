package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	greeting := greet("Pradeep")
	if !strings.Contains(greeting, "Hey ") {
		t.Error("Response from greet is unexpected")
	}
}
