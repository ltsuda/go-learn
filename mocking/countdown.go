package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownFrom = 3

func Countdown(output io.Writer) {
	for i := countdownFrom; i > 0; i-- {
		fmt.Fprintln(output, i) //nolint:errcheck
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(output, finalWorld) //nolint:errcheck
}

func main() {
	Countdown(os.Stdout)
}
