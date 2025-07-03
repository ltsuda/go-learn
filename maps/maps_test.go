package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is just a test"
		assertStrings(t, result, expected)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	result, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, result, expected)
}

func assertStrings(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}

func assertError(t testing.TB, result, expected error) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %q got %q", expected.Error(), result)
	}
}
