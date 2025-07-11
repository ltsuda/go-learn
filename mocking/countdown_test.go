package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	result := buffer.String()
	expected := `3
2
1
Go!`

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, expected 3 but got %d", spySleeper.Calls)
	}

}
