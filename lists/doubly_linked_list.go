package lists

import "fmt"

type doublyNode[T any] struct {
	value T
	next  *doublyNode[T]
	prev  *doublyNode[T]
}

type DoublyLinkedList[T any] struct {
	Length     int
	head, tail *doublyNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head:   nil,
		tail:   nil,
		Length: 0,
	}
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.Length == 0
}

func (l *DoublyLinkedList[T]) Append(value T) *DoublyLinkedList[T] {
	node := &doublyNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if l.IsEmpty() {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.Length += 1
	l.tail = node

	return l
}

func (l *DoublyLinkedList[T]) Preppend(value T) *DoublyLinkedList[T] {
	node := &doublyNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if l.IsEmpty() {
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
	}

	l.Length += 1
	l.head = node

	return l
}

func (l *DoublyLinkedList[T]) InsertAt(value T, idx int) *DoublyLinkedList[T] {
	if idx == 0 {
		return l.Preppend(value)
	}

	if idx == l.Length {
		return l.Append(value)
	}

	current := l.head
	previous := l.head
	for i := 0; i < idx; i++ {
		previous = current
		current = current.next
	}

	node := &doublyNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	previous.next = node
	node.next = current
	node.prev = previous
	current.prev = node
	l.Length += 1

	return l
}

func (l *DoublyLinkedList[T]) Get(idx int) (T, error) {
	if idx == 0 {
		return l.head.value, nil
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

func (l *DoublyLinkedList[T]) Remove(idx int) (T, error) {
	if idx == 0 {
		node := l.head

		l.head = l.head.next
		l.head.prev = nil
		node.next = nil

		return node.value, nil
	}

	if idx == l.Length-1 {
		node := l.tail

		l.tail = l.tail.prev
		l.tail.next = nil
		node.prev = nil

		return node.value, nil
	}

	prev := l.head
	current := l.head
	for i := 0; i < idx; i++ {
		prev = current
		current = current.next
	}

	prev.next = current.next
	current.next.prev = prev
	current.next = nil
	current.prev = nil
	l.Length -= 1

	return current.value, nil
}

func (l *DoublyLinkedList[T]) PrintAllElements() {
	current := l.head
	for i := 0; i < l.Length; i++ {
		fmt.Println(current.value)
		current = current.next
	}
}

func (l *DoublyLinkedList[T]) PrintAllReversedElements() {
	current := l.tail
	for i := l.Length; i > 0; i-- {
		fmt.Println(current.value)
		current = current.prev
	}
}
