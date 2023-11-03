package queue

import "fmt"

type node[T any] struct {
	value T
	prev  *node[T]
}

type Queue[T any] struct {
	length     int
	head, tail *node[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Enqueue(value T) *Queue[T] {
	node := &node[T]{
		value: value,
		prev:  nil,
	}

	if q.length == 0 {
		q.head = node
	} else {
		q.tail.prev = node
	}

	q.tail = node
	q.length += 1

	return q
}

func (q *Queue[T]) Dequeue() T {
	v := q.head.value
	node := q.head

	q.head = q.head.prev
	node.prev = nil
	q.length -= 1

	return v
}

func (q *Queue[T]) PrintQueue() {
	current := q.head
	for i := 0; i < q.length; i++ {
		fmt.Println(i+1, current.value)
		current = current.prev
	}
}
