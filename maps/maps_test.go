package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	result := dictionary.Search("test")
	expected := "this is just a test"

	assertStrings(t, result, expected)
}

func assertStrings(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}
