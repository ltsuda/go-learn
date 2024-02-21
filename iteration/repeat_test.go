package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("z", 10)
	expected := "zzzzzzzzzz"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 5)
	}
}

func ExampleRepeat() {
	repeatChar := Repeat("a", 5)
	fmt.Println(repeatChar)
	// Output: aaaaa
}
