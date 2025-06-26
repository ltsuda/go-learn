package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := Perimeter(rectangle)
	expected := 40.0

	if perimeter != expected {
		t.Errorf("expected %.2f, got %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		area := shape.Area()

		if area != expected {
			t.Errorf("expected %g, got %g", expected, area)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 12.0}
		checkArea(t, rectangle, 120.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
