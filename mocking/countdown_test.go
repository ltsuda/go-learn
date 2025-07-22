package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)

		result := buffer.String()
		expected := `3
2
1
Go!`

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spySleepPrinter.Calls) {
			t.Errorf("expected calls %v, got %v", expected, spySleepPrinter.Calls)
		}
	})

	t.Run("countDown with iterator", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		CountdownWithIterator(buffer, spySleeper)

		result := buffer.String()
		expected := `3
2
1
Go!`

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("sleep for 5", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}

	})

}
