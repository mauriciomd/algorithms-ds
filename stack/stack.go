package stack

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	length int
	head   *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		head:   nil,
		length: 0,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Push(value T) *Stack[T] {
	node := &node[T]{
		value: value,
		next:  nil,
	}

	if !s.IsEmpty() {
		node.next = s.head
	}

	s.length += 1
	s.head = node

	return s
}

func (s *Stack[T]) Pop() T {
	node := s.head

	s.head = s.head.next
	node.next = nil

	return node.value
}

func (s Stack[T]) PrintStack() {
	current := s.head
	for i := 0; i < s.length; i++ {
		fmt.Println(i+1, current.value)
		current = current.next
	}
}
