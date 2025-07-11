package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownFrom = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(output io.Writer, sleeper Sleeper) {
	for i := countdownFrom; i > 0; i-- {
		fmt.Fprintln(output, i) //nolint:errcheck
		sleeper.Sleep()
	}
	fmt.Fprint(output, finalWorld) //nolint:errcheck
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
