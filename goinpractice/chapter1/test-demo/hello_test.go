package main

import "testing"

func TestName(t *testing.T) {
	name := getName()

	if name != "Hello World!" {
		t.Error("Response from getName() is unexpected value!!")
	}
}
