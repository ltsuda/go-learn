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
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	result, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, result, definition)
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
