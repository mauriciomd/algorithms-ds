package tree

type MinHeap struct {
	heap   []int
	length int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap:   make([]int, 16),
		length: 0,
	}
}

func getParent(idx int) int {
	return (idx - 1) / 2
}

func getLeftChild(idx int) int {
	return idx*2 + 1
}

func getRightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parent := getParent(idx)
	if h.heap[idx] < h.heap[parent] {
		h.heap[parent], h.heap[idx] = h.heap[idx], h.heap[parent]
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	leftIdx := getLeftChild(idx)
	rightIdx := getRightChild(idx)

	if idx >= h.length || leftIdx >= h.length {
		return
	}

	if h.heap[leftIdx] > h.heap[rightIdx] && h.heap[idx] > h.heap[rightIdx] {
		h.heap[idx], h.heap[rightIdx] = h.heap[rightIdx], h.heap[idx]
		h.heapifyDown(rightIdx)
	} else if h.heap[rightIdx] > h.heap[leftIdx] && h.heap[idx] > h.heap[leftIdx] {
		h.heap[idx], h.heap[leftIdx] = h.heap[leftIdx], h.heap[idx]
		h.heapifyDown(leftIdx)
	}
}

func (h *MinHeap) Insert(value int) {
	h.heap[h.length] = value
	h.heapifyUp(h.length)
	h.length += 1
}

func (h *MinHeap) Poll() int {
	if h.length == 0 {
		return -1
	}

	out := h.heap[0]
	h.length -= 1

	if h.length == 0 {
		return out
	}

	h.heap[0] = h.heap[h.length]
	h.heapifyDown(0)

	return out
}
