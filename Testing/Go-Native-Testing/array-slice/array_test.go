package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 Numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		response := Sum(numbers)
		want := 15

		if response != want {
			t.Errorf("got %d want %d given %v", response, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{9, 8, 7, 6, 5})
	want := []int{6, 35}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
