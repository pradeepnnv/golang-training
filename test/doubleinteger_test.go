package mytest

import "testing"

func TestDoubleInteger(t *testing.T) {
	if dI := doubleTheInteger(2); dI != 4 {
		t.Error("Unexpected response")
	}
}
