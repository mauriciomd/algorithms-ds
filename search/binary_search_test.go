package search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 11, 23}

	t.Run("Element exists in the middle", func(t *testing.T) {
		got := BinarySearch(7, arr)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Element is in the first position", func(t *testing.T) {
		got := BinarySearch(1, arr)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Element doesn't exist", func(t *testing.T) {
		got := BinarySearch(14, arr)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
