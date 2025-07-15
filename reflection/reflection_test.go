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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var result []string
		walk(aMap, func(input string) {
			result = append(result, input)
		})

		assertContains(t, result, "Moo")
		assertContains(t, result, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{34, "Calgary"}
			close(aChannel)
		}()

		var result []string
		expected := []string{"London", "Calgary"}

		walk(aChannel, func(input string) {
			result = append(result, input)
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, result %v", expected, result)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{40, "London"}, Profile{50, "Calgary"}
		}
		var result []string
		expected := []string{"London", "Calgary"}

		walk(aFunction, func(input string) {
			result = append(result, input)
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, result %v", expected, result)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
