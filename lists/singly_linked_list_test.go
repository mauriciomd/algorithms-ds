package lists

import (
	"errors"
	"testing"
)

func TestAppendSinglyLinkedList(t *testing.T) {
	t.Run("Trying to insert the first element", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(10)
		got := lst.tail
		want := lst.head

		if got != want || want.value != 10 {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Trying to insert three elements", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(10)
		lst.Append(20)
		lst.Append(30)

		got := lst.tail.value
		want := 30

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestPreppendSinglyLinkedList(t *testing.T) {
	t.Run("Trying to preppend the first element", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Preppend(100)
		got := lst.tail
		want := lst.head

		if got != want || want.value != 100 {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Trying to preppend three elements", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Preppend(10)
		lst.Preppend(20)
		lst.Preppend(30)

		got := lst.head.value
		want := 30

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetSinglyLinkedList(t *testing.T) {
	t.Run("Trying to get an element from an empty list", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		_, got := lst.Get(1)
		want := errors.New("Empty list")

		if got.Error() != want.Error() {
			t.Errorf("got %s, want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to get an element from index out of range", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(1)
		lst.Append(2)

		_, got := lst.Get(2)
		want := errors.New("Index out of range")

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to get an element from negative index", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(1)
		lst.Append(2)

		_, got := lst.Get(-1)
		want := errors.New("Index cannot be negative")

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to get a specific element", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(1)
		lst.Append(2)
		lst.Append(3)
		lst.Append(4)

		got, _ := lst.Get(0)
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		got, _ = lst.Get(2)
		want = 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		got, _ = lst.Get(3)
		want = 4
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestRemoveSinglyLinkedList(t *testing.T) {
	t.Run("Trying to remove an element from empty list", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		_, got := lst.Remove(1)
		want := errors.New("Empty list")

		if got.Error() != want.Error() {
			t.Errorf("got %s, want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to remove from index out of range", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(10)

		_, got := lst.Remove(2)
		want := errors.New("Index out of range")

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to remove an element from negative index", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()

		lst.Append(1)
		lst.Append(2)

		_, got := lst.Get(-1)
		want := errors.New("Index cannot be negative")

		if got.Error() != want.Error() {
			t.Errorf("got %s want %s", got.Error(), want.Error())
		}
	})

	t.Run("Trying to remove the first element", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()
		lst.Append(1)
		lst.Append(2)

		lst.Remove(0)

		got, _ := lst.Get(0)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Trying to remove an element in the middle", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()
		lst.Append(10)
		lst.Append(20)
		lst.Append(30)
		lst.Append(40)
		lst.Append(50)

		lst.Remove(2)

		got, _ := lst.Get(2)
		want := 40

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Trying to remove the last element", func(t *testing.T) {
		lst := NewSinglyLinkedList[int]()
		lst.Append(9)
		lst.Append(8)
		lst.Append(7)

		lst.Remove(2)

		got := lst.tail.value
		want := 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
