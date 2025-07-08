package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mike")

	result := buffer.String()
	expected := "Hello, Mike"

	if result != expected {
		t.Errorf("expected %q, result %q", expected, result)
	}
}
