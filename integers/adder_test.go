package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 2)
	expected := 6
	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 9)
	fmt.Println(sum)
	//Output: 12
}
