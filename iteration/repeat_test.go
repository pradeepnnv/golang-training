package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 3)
	expected := "aaa"
	if expected != actual {
		t.Errorf("Expected %q. But got %q.", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("aadfasdfsdfadsfsadfasdfasdfadsfasdfasd", 90000)
	}
}

func ExampleRepeat() {
	o := Repeat("a", 5)
	fmt.Println(o)
	//Output: aaaaa
}
