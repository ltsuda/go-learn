package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it as 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, result Counter, expected int) {
	t.Helper()
	if result.Value() != expected {
		t.Errorf("expected %d, got %d", result.Value(), 3)
	}
}
