package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("expected %d but got %d, %v", expected, sum, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	sum := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v but got %v", expected, result)
		}
	}

	t.Run("make the sums o f some slices", func(t *testing.T) {

		sum := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, sum, expected)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, sum, expected)
	})
}
