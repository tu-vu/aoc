package stack

import (
	"errors"
)

type Stack[T any] []T

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop removes and returns a value from the top of the stack.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

// Peek returns the top value from the stack without removing it.
func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}
