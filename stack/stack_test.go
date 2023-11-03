package stack

import "testing"

func TestPush(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	got := s.head.value
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	got := s.Pop()
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = s.Pop()
	want = 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = s.Pop()
	want = 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
