package lists

import "errors"

func (l *SinglyLinkedList[T]) formatError(msg string) (T, error) {
	return *new(T), errors.New(msg)
}
