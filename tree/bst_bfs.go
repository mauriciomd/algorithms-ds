package tree

import (
	"github.com/mauriciomd/algorithms-ds/queue"
)

func BSF(t *BSTree) []int {
	q := queue.New[*node]()
	elements := []int{}

	if t.root == nil {
		return elements
	}

	q.Enqueue(t.root)

	for !q.IsEmpty() {
		node := q.Dequeue()
		elements = append(elements, node.data)

		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}
	}

	return elements
}
