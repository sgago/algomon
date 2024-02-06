// Package stack implements a first-in last-out (FILO) collection.
package stack

import (
	"fmt"
	"slices"

	"github.com/sgago/algomon/col/conc"
	"github.com/sgago/algomon/errs"
)

// Stack is a a first-in last-out (FILO) collection.
type Stack[T any] struct {
	conc.RWLocker

	// The elements in this collection.
	v []T
}

// New creates a new Stack and initializes it with the given values.
func New[T any](cap int, vals ...T) *Stack[T] {
	cap = max(cap, len(vals))

	v := make([]T, len(vals))
	copy(v, vals)
	slices.Reverse(v)

	return &Stack[T]{
		v: append(make([]T, 0, cap), v...),
	}
}

// Slice returns the internal slice backing this collection.
func (s *Stack[T]) Slice() []T {
	defer s.RWLocker.Read()()

	return s.v
}

// String returns this collection represented as a string.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

// Len returns the number of elements in this Stack.
func (s *Stack[T]) Len() int {
	defer s.RWLocker.Read()()

	return len(s.v)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Cap returns the maximum number of elements this Stack can have before being re-sized.
func (s *Stack[T]) Cap() int {
	defer s.RWLocker.Read()()

	return cap(s.v)
}

// Push adds values to the top of the Stack.
func (s *Stack[T]) Push(vals ...T) {
	defer s.RWLocker.Write()()

	v := make([]T, len(vals))
	copy(v, vals)
	slices.Reverse(v)

	s.v = append(v, s.v...)
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack[T]) Peek() T {
	defer s.RWLocker.Read()()
	return s.v[0]
}

// TryPeek attempts to peek at the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPeek() (T, error) {
	defer s.RWLocker.Read()()

	if len(s.v) > 0 {
		return s.v[0], nil
	}

	return *new(T), errs.Empty
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack[T]) Pop() T {
	defer s.RWLocker.Write()()

	result := s.v[0]
	s.v = s.v[:len(s.v)-1]
	return result
}

// TryPop attempts to pop a value from the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPop() (T, error) {
	defer s.RWLocker.Write()()

	if len(s.v) > 0 {
		result := s.v[0]
		s.v = s.v[:len(s.v)-1]
		return result, nil
	}

	return *new(T), errs.Empty
}
