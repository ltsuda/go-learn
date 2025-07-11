package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownFrom = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(output io.Writer, sleeper Sleeper) {
	for i := countdownFrom; i > 0; i-- {
		fmt.Fprintln(output, i) //nolint:errcheck
		sleeper.Sleep()
	}
	fmt.Fprint(output, finalWorld) //nolint:errcheck
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
