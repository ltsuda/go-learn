package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		output := Hello("John Doe", "")
		expected := "Hello, John Doe"

		assertCorrectMessage(t, output, expected)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		output := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, output, expected)
	})
	t.Run("in Spanish", func(t *testing.T) {
		output := Hello("John Spain", "Spanish")
		expected := "Hola, John Spain"
		assertCorrectMessage(t, output, expected)
	})
	t.Run("in French", func(t *testing.T) {
		output := Hello("John French", "French")
		expected := "Bonjour, John French"
		assertCorrectMessage(t, output, expected)
	})
}

func assertCorrectMessage(t testing.TB, output, expected string) {
	t.Helper()
	if output != expected {
		t.Errorf("output %q expected %q", output, expected)
	}
}
