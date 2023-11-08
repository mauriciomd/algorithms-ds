package tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestBfs(t *testing.T) {
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

	t.Run("Breadth First Traversal", func(t *testing.T) {
		want := "8,3,10,1,6,14,4,7,13"
		got := strings.Join(Map(BSF(bst), strconv.Itoa), ",")

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
