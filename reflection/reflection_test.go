package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Mike"
	var result []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		result = append(result, input)
	})

	if len(result) != 1 {
		t.Errorf("wrong number of function calls, result %d, expected %d", len(result), 1)
	}

	if result[0] != expected {
		t.Errorf("expected %q, got %q", expected, result[0])
	}
}
