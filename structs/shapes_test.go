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

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 12.0}, expected: 120},
		{name: "Circle", shape: Circle{Radius: 10.0}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 8}, expected: 40.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			area := tt.shape.Area()

			if area != tt.expected {
				t.Errorf("expected %g, got %g", tt.expected, area)
			}
		})
	}
}
