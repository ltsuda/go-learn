package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a) //nolint:errcheck
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b) //nolint:errcheck
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}
