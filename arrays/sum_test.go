package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		got := Sum(numbers)
		want := 9
		if got != want {
			t.Errorf("want %d but got %d,given %v", want, got, numbers)
		}
	})
}
