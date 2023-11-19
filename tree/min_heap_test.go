package tree

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	heap.Insert(11)
	heap.Insert(14)
	heap.Insert(10)
	heap.Insert(21)
	heap.Insert(13)
	heap.Insert(9)
	heap.Insert(19)
	// Expected array after insertions
	// 9, 13, 10, 21, 14, 11, 19

	t.Run("Try to insert", func(t *testing.T) {
		got := heap.length
		want := 7

		if want != got {
			t.Errorf("got %d want %d", got, want)
		}

		got = heap.heap[0]
		want = 9
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}

		got = heap.heap[2]
		want = 10
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Try to poll", func(t *testing.T) {
		got := heap.Poll()
		want := 9
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}

		// Expected array after poll
		// 10, 13, 11, 21, 14, 19
		got = heap.heap[2]
		want = 11

		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
