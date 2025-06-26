package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeter := Perimeter(10.0, 10.0)
	expected := 40.0

	if perimeter != expected {
		t.Errorf("expected %.2f, got %.2f", expected, perimeter)
	}
}

func TestArea(t *testing.T) {
	area := Area(10.0, 12.0)
	expected := 120.0

	if area != expected {
		t.Errorf("expected %.2f, got %.2f", expected, area)
	}
}
