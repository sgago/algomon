// Package stack implements a first-in last-out (FILO) collection.
package stack

import (
	"fmt"
	"slices"
	"sync"

	"github.com/sgago/algomon/errs"
)

// Stack is a a first-in last-out (FILO) collection.
type Stack[T any] struct {
	// The elements in this collection.
	v []T
	// Indicates if this is a synchronized collection.
	sync bool
	// Mutex for synchronizing this collection.
	mu sync.RWMutex
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
	if s.IsSynchronized() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	return s.v
}

// String returns this collection represented as a string.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

// Len returns the number of elements in this Stack.
func (s *Stack[T]) Len() int {
	if s.IsSynchronized() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	return len(s.v)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Cap returns the maximum number of elements this Stack can have before being re-sized.
func (s *Stack[T]) Cap() int {
	if s.IsSynchronized() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	return cap(s.v)
}

// Push adds values to the top of the Stack.
func (s *Stack[T]) Push(vals ...T) {
	if s.IsSynchronized() {
		defer s.mu.Unlock()
		s.mu.Lock()
	}

	v := make([]T, len(vals))
	copy(v, vals)
	slices.Reverse(v)

	s.v = append(v, s.v...)
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack[T]) Peek() T {
	if s.IsSynchronized() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	return s.v[0]
}

// TryPeek attempts to peek at the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPeek() (T, error) {
	if s.IsSynchronized() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	if len(s.v) > 0 {
		return s.v[0], nil
	}

	return *new(T), errs.Empty
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack[T]) Pop() T {
	if s.IsSynchronized() {
		defer s.mu.Unlock()
		s.mu.Lock()
	}

	result := s.v[0]
	s.v = s.v[:len(s.v)-1]
	return result
}

// TryPop attempts to pop a value from the top of the stack and returns an error if the stack is empty.
func (s *Stack[T]) TryPop() (T, error) {
	if s.IsSynchronized() {
		defer s.mu.Unlock()
		s.mu.Lock()
	}

	if len(s.v) > 0 {
		result := s.v[0]
		s.v = s.v[:len(s.v)-1]
		return result, nil
	}

	return *new(T), errs.Empty
}

// Synchronize sets the synchronization flag for the stack.
// If enable is true, it enables synchronization, meaning that
// concurrent access to the stack will be synchronized using locks.
// If enable is false, it disables synchronization, allowing
// concurrent access without additional synchronization.
func (s *Stack[T]) Synchronize(enable bool) {
	s.sync = enable
}

// IsSynchronized returns true if the stack is configured for synchronization,
// meaning that concurrent access is synchronized using locks.
// Returns false if synchronization is disabled,
// allowing concurrent access without additional synchronization.
func (s *Stack[T]) IsSynchronized() bool {
	return s.sync
}
