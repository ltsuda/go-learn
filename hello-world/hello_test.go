package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		output := Hello("John Doe")
		expected := "Hello, John Doe"

		if output != expected {
			t.Errorf("output %q expected %q", output, expected)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		output := Hello("")
		expected := "Hello, World"

		if output != expected {
			t.Errorf("output %q expected %q", output, expected)
		}
	})
}
