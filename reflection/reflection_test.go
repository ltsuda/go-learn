package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

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
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Mike", 34},
			[]string{"Mike"},
		},
		{
			"struct with nested fields",
			Person{
				"Mike",
				Profile{34, "London"},
			},
			[]string{"Mike", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Mike",
				Profile{34, "London"},
			},
			[]string{"Mike", "London"},
		},
		{
			"slices",
			[]Profile{
				{34, "London"},
				{40, "Calgary"},
			},
			[]string{"London", "Calgary"},
		},
		{
			"arrays",
			[2]Profile{
				{40, "London"},
				{50, "Calgary"},
			},
			[]string{"London", "Calgary"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
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
