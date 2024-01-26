// Package casstack implements a compare-and-swap (CAS) first-in
// last-out (FILO) collection that supports concurrent operations.
// Reference: https://www.geeksforgeeks.org/lock-free-stack-using-java/
package casstack

import (
	"sync/atomic"

	"github.com/sgago/algomon/errs"
)

// node represents a node in the CAS stack, containing a value and a pointer to the next node.
type node[T any] struct {
	// The node's value.
	val T

	// A pointer to the next node.
	next atomic.Pointer[node[T]]
}

// CasStack is a compare-and-swap (CAS) stack that supports concurrent operations.
type CasStack[T any] struct {
	// A pointer to the top (head) node of the stack.
	head atomic.Pointer[node[T]]

	// The number of elements in the stack.
	len atomic.Int32
}

// New creates a new CasStack and initializes it with the given values.
func New[T any](vals ...T) *CasStack[T] {
	s := &CasStack[T]{}

	for _, v := range vals {
		s.Push(v)
	}

	return s
}

// Len returns the current length of the stack.
func (s *CasStack[T]) Len() int {
	return int(s.len.Load())
}

// Push adds values to the top of the stack using CAS operations for concurrency safety.
func (s *CasStack[T]) Push(vals ...T) {
	for _, v := range vals {
		new := &node[T]{
			val:  v,
			next: s.head,
		}

		for !s.head.CompareAndSwap(s.head.Load(), new) {
			new.next = s.head
		}

		s.len.Add(1)
	}
}

// Peek returns the value at the top of the stack without removing it.
func (s *CasStack[T]) Peek() T {
	return s.head.Load().val
}

// TryPeek attempts to peek at the top of the stack and returns an error if the stack is empty.
func (s *CasStack[T]) TryPeek() (T, error) {
	h := s.head.Load()

	if h == nil {
		return *new(T), errs.Empty
	}

	s.len.Add(-1)

	return h.val, nil
}

// Pop removes and returns the value at the top of the stack.
func (s *CasStack[T]) Pop() T {
	curr := s.head.Load()

	// This is where the magic happens.
	// If s.head is the same as curr, then we swap.
	// If s.head is different because another thread modified it, we spin, reload curr, and try again.
	for !s.head.CompareAndSwap(curr, curr.next.Load()) {
		curr = s.head.Load()
	}

	s.len.Add(-1)

	return curr.val
}

// TryPop attempts to pop a value from the top of the stack and returns an error if the stack is empty.
func (s *CasStack[T]) TryPop() (T, error) {
	curr := s.head.Load()

	for curr != nil && !s.head.CompareAndSwap(curr, curr.next.Load()) {
		curr = s.head.Load()
	}

	if curr == nil {
		return *new(T), errs.Empty
	}

	s.len.Add(-1)

	return curr.val, nil
}
