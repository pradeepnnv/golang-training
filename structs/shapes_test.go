package structs

import (
	"testing"
)

// func TestPerimeter(t *testing.T) {

// 	t.Run("calculate perimeter of a rectangle", func(t *testing.T) {
// 		r := Rectangle{10.0, 2.0}
// 		got := r.Perimeter()
// 		want := 24.0

// 		if got != want {
// 			t.Errorf("got %.2f but want %.2f", got, want)
// 		}
// 	})

// 	t.Run("calculate perimeter of a circle", func(t *testing.T) {
// 		c := Circle{10.0}
// 		got := c.Perimeter()
// 		want := (2 * math.Pi * 10.0)

// 		if got != want {
// 			t.Errorf("got %.2f but want %.2f", got, want)
// 		}
// 	})
// }

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("calculate area of a rectangle", func(t *testing.T) {
		checkArea(t, Rectangle{10.0, 2.0}, 20.0)
	})

	t.Run("calculate area of a circle", func(t *testing.T) {
		checkArea(t, Circle{1}, 3.141592653589793)
	})
}
