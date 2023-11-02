package lists

import "testing"

func TestAppendDoublyLinkedList(t *testing.T) {
	t.Run("Trying to append the first element", func(t *testing.T) {
		lst := NewDoublyLinkedList[int]()
		lst.Append(21)

		got := lst.head.value
		want := lst.tail.value

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("Trying to append elements", func(t *testing.T) {
		lst := NewDoublyLinkedList[int]()
		lst.Append(21)
		lst.Append(32)
		lst.Append(11)

		got := lst.tail.value
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func TestPreppendDoublyLinkedList(t *testing.T) {
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

		lst.Preppend(40)
		lst.Preppend(60)
		lst.Preppend(10)

		got := lst.head.value
		want := 10

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetDoublyLinkedList(t *testing.T) {
	lst := NewDoublyLinkedList[int]()

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
}

func TestInsertAtDoublyLinkedList(t *testing.T) {
	lst := NewDoublyLinkedList[int]()

	lst.Append(1)
	lst.Append(2)
	lst.Append(3)

	lst.InsertAt(4, 1)

	got, _ := lst.Get(1)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestRemoveDoublyLinkedList(t *testing.T) {
	t.Run("Trying to remove the first element", func(t *testing.T) {
		lst := NewDoublyLinkedList[int]()
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
		lst := NewDoublyLinkedList[int]()
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
		lst := NewDoublyLinkedList[int]()
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
