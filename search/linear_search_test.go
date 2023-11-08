package search

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 4, 6, 21, 5, 11, 17}

	t.Run("Element exists", func(t *testing.T) {
		got := LinearSearch(11, arr)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Element doesn't exist", func(t *testing.T) {
		got := LinearSearch(99, arr)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
