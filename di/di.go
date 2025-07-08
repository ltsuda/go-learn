package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name) //nolint:errcheck
}

// func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

// func di() {
// 	Greet(os.Stdout, "Mike")
// }

// func main() {
// 	log.Fatal(http.ListenAndServe(":5005", http.HandlerFunc(MyGreetHandler)))
// }
