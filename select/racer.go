package racer

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url) //nolint:errcheck
	return time.Since(start)
}

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}
