package car

import (
	"testing"
)

func TestNew(t *testing.T) {
	want := "BMW Gray (model)"
	if got := New(BMW, Gray, "model"); got.String() != want {
		t.Errorf("New() = %q, want %q", got, want)
	}
}
