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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 12.0}
		area := rectangle.Area()
		expected := 120.0

		if area != expected {
			t.Errorf("expected %.2f, got %.2f", expected, area)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		area := circle.Area()
		expected := 314.1592653589793

		if area != expected {
			t.Errorf("expected %g, got %g", expected, area)
		}
	})
}
