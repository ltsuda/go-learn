package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	result := Search(dictionary, "test")
	expected := "this is just a test"

	if result != expected {
		t.Errorf("expected %q got %q given, %q", expected, result, "test")
	}
}
