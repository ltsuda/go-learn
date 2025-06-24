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
	sum := SumAllTails([]int{1, 2}, []int{0, 9})
	expected := []int{2, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}
}
