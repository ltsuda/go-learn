package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Mike"},
			[]string{"Mike"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mike", "London"},
			[]string{"Mike", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string
			walk(test.Input, func(input string) {
				result = append(result, input)
			})

			if !reflect.DeepEqual(result, test.ExpectedCalls) {
				t.Errorf("expected %v, result %v", test.ExpectedCalls, result)
			}
		})
	}
}
