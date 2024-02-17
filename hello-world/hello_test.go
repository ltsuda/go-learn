package main

import "testing"

func TestHello(t *testing.T) {
	output := Hello("John Doe")
	expected := "Hello, John Doe"

	if output != expected {
		t.Errorf("output %q expected %q", output, expected)
	}
}
