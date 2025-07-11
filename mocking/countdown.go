package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(output io.Writer) {
	fmt.Fprint(output, "3") //nolint:errcheck
}

func main() {
	Countdown(os.Stdout)
}
