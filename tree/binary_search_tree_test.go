package tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestInsertBST(t *testing.T) {
	bst := New()

	bst.Insert(8)
	bst.Insert(3)
	bst.Insert(10)

	got := bst.root.data
	want := 8

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = bst.root.left.data
	want = 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = bst.root.right.data
	want = 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/*
			      8
		 	   /    \
		     /       \
           /          \
         /             \
	   3               10
	 /  \                \
	/    \	              \
   1      6               14
	     / \            /
	    4   7         13

*/

func TestPrintBST(t *testing.T) {
	bst := New()

	bst.Insert(8)
	bst.Insert(3)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(14)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(13)

	t.Run("Pre order", func(t *testing.T) {
		want := "8,3,1,6,4,7,10,14,13"
		got := strings.Join(Map(bst.GetElementsPreOrder(), strconv.Itoa), ",")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("In order", func(t *testing.T) {
		want := "1,3,4,6,7,8,10,13,14"
		got := strings.Join(Map(bst.GetElementsInOrder(), strconv.Itoa), ",")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Pos order", func(t *testing.T) {
		want := "1,4,7,6,3,13,14,10,8"
		got := strings.Join(Map(bst.GetElementsPosOrder(), strconv.Itoa), ",")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
