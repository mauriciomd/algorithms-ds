package queue

import "testing"

func TestEnqueue(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	got := queue.tail.value
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDequeue(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	got := queue.Dequeue()
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = queue.Dequeue()
	want = 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = queue.Dequeue()
	want = 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
