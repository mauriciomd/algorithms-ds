package lists

import (
	"fmt"
)

type singlyNode[T any] struct {
	value T
	next  *singlyNode[T]
}

type SinglyLinkedList[T any] struct {
	Length     int
	head, tail *singlyNode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.Length == 0
}

func (l *SinglyLinkedList[T]) Append(value T) *SinglyLinkedList[T] {
	node := &singlyNode[T]{
		value: value,
		next:  nil,
	}

	if l.IsEmpty() {
		l.head = node
	} else {
		l.tail.next = node
	}

	l.tail = node
	l.Length += 1

	return l
}

func (l *SinglyLinkedList[T]) Preppend(value T) *SinglyLinkedList[T] {
	node := &singlyNode[T]{
		value: value,
		next:  l.head,
	}

	if l.Length == 0 {
		l.tail = node
	}

	l.Length += 1
	l.head = node

	return l
}

func (l *SinglyLinkedList[T]) Get(idx int) (T, error) {
	if l.Length == 0 {
		return l.formatError("Empty list")
	}

	if idx >= l.Length {
		return l.formatError("Index out of range")
	}

	if idx < 0 {
		return l.formatError("Index cannot be negative")
	}

	if idx == l.Length-1 {
		return l.tail.value, nil
	}

	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}

	return current.value, nil
}

func (l *SinglyLinkedList[T]) Remove(idx int) (T, error) {
	if l.IsEmpty() {
		return l.formatError("Empty list")
	}

	if idx >= l.Length {
		return l.formatError("Index out of range")
	}

	if idx < 0 {
		return l.formatError("Index cannot be negative")
	}

	if idx == 0 {
		value := l.head.value
		l.head = l.head.next

		return value, nil
	}

	current := l.head
	previous := l.head
	for i := 0; i < idx; i++ {
		previous = current
		current = current.next
	}

	value := current.value
	previous.next = current.next

	if idx == l.Length-1 {
		l.tail = previous
	}

	l.Length -= 1

	return value, nil
}

func (l *SinglyLinkedList[T]) PrintAllElements() {
	current := l.head
	for i := 0; i < l.Length; i++ {
		fmt.Println(current.value)
		current = current.next
	}
}
