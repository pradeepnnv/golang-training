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

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Length: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10.0, 2.0}, hasArea: 10},
	}
	for _, tt := range areaTests {
		// got := tt.shape.Area()
		// if got != tt.hasArea {
		// 	t.Errorf("%#v got %g but want %g", tt.shape, got, tt.hasArea)
		// }
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g but want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
