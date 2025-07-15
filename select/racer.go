package racer

import (
	"net/http"
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url) //nolint:errcheck
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
