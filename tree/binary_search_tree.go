package tree

type node struct {
	data        int
	left, right *node
}

type BSTree struct {
	root *node
}

func New() *BSTree {
	return &BSTree{
		root: nil,
	}
}

func (t *BSTree) Insert(value int) {
	t.root = insert(t.root, value)
}

func insert(n *node, value int) *node {
	if n == nil {
		return &node{
			data:  value,
			left:  nil,
			right: nil,
		}
	} else if n.data > value {
		n.left = insert(n.left, value)
	} else {
		n.right = insert(n.right, value)
	}

	return n
}

func (t *BSTree) GetElementsPreOrder() []int {
	path := []int{}
	preOrder(t.root, &path)

	return path
}

func preOrder(n *node, path *[]int) {
	if n == nil {
		return
	}

	*path = append(*path, n.data)
	preOrder(n.left, path)
	preOrder(n.right, path)
}

func (t *BSTree) GetElementsInOrder() []int {
	path := []int{}
	inOrder(t.root, &path)

	return path
}

func inOrder(n *node, path *[]int) {
	if n == nil {
		return
	}

	inOrder(n.left, path)
	*path = append(*path, n.data)
	inOrder(n.right, path)
}

func (t *BSTree) GetElementsPosOrder() []int {
	path := []int{}
	posOrder(t.root, &path)

	return path
}

func posOrder(n *node, path *[]int) {
	if n == nil {
		return
	}

	posOrder(n.left, path)
	posOrder(n.right, path)
	*path = append(*path, n.data)
}
