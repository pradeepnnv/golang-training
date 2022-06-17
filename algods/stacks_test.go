package algods

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	var stack Stack
	stack.Push('a')
	stack.Push('b')

	if len(stack) != 2 {
		t.Error("Elements are not pushed into the stack")
	}
}

func TestStackPop(t *testing.T) {
	var stack Stack
	stack.Push('c')
	stack.Push('d')

	val, ok := stack.Pop()
	if !ok || val != 'd' {
		t.Error("Element is not popped successfully")
	}
}

func TestIsValid(t *testing.T) {
	s := "()"
	expected := true

	if isValid(s) != expected {
		t.Error("IsValid is wrong")
	}
}
