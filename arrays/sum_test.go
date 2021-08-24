package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3}, []int{4, 5, 6, 7, 8})
	want := []int{5, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{4, 5, 6, 7, 8})
		want := []int{3, 26}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{2, 3}, []int{4, 5, 6, 7, 8})
	}
}
