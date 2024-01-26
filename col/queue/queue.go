// Package queue implements a first-in first-out (FIFO) collection.
package queue

import (
	"fmt"
	"sync"

	"github.com/sgago/algomon/errs"
)

// Queue is a a first-in first-out (FIFO) collection.
type Queue[T any] struct {
	// The elements in this collection.
	v []T
	// Indicates if this is a synchronized collection.
	sync bool
	// Mutex for synchronizing this collection.
	mu sync.RWMutex
}

// New creates a new Queue and initializes it with the given values.
func New[T any](cap int, vals ...T) *Queue[T] {
	cap = max(cap, len(vals))

	return &Queue[T]{
		v: append(make([]T, 0, cap), vals...),
	}
}

// Slice returns the internal slice backing this collection.
func (s *Queue[T]) Slice() []T {
	if s.IsSync() {
		defer s.mu.RUnlock()
		s.mu.RLock()
	}

	return s.v
}

// String returns this collection represented as a string.
func (q *Queue[T]) String() string {
	if q.IsSync() {
		defer q.mu.RUnlock()
		q.mu.RLock()
	}

	return fmt.Sprintf("%v", q.v)
}

// Len returns the number of elements in this Queue.
func (q *Queue[T]) Len() int {
	if q.IsSync() {
		defer q.mu.RUnlock()
		q.mu.RLock()
	}

	return len(q.v)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

// Cap returns the maximum number of elements this Queue can have before being re-sized.
func (q *Queue[T]) Cap() int {
	if q.IsSync() {
		defer q.mu.RUnlock()
		q.mu.RLock()
	}

	return cap(q.v)
}

// Enq adds values to the Queue.
func (q *Queue[T]) Enq(vals ...T) *EnqOp[T] {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	return &EnqOp[T]{
		v: vals,
		q: q,
	}
}

// Deq removes and returns the value at the end of the Queue.
func (q *Queue[T]) Deq() *DeqOp[T] {
	if q.IsSync() {
		defer q.mu.Unlock()
		q.mu.Lock()
	}

	return &DeqOp[T]{
		q: q,
	}
}

// TryDeq attempts to pop a value from the start of the Queue and returns an error if the Queue is empty.
func (q *Queue[T]) TryDeq() (T, error) {
	if q.IsSync() {
		defer q.mu.Unlock()
		q.mu.Lock()
	}

	if len(q.v) > 0 {
		result := q.v[0]
		q.v = q.v[:len(q.v)-1]
		return result, nil
	}

	return *new(T), errs.Empty
}

// Peek returns the value at the start of the Queue without removing it.
func (q *Queue[T]) Peek() *PeekOp[T] {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	return &PeekOp[T]{
		q: q,
	}
}

// TryPeek attempts to peek at the start of the Queue and returns an error if the Queue is empty.
func (q *Queue[T]) TryPeek() (T, error) {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	if len(q.v) > 0 {
		return q.v[0], nil
	}

	return *new(T), errs.Empty
}

// Sync sets the synchronization flag for the collection.
// If enable is true, it enables synchronization, meaning that
// concurrent access to the Queue will be synchronized using locks.
// If enable is false, it disables synchronization, allowing
// concurrent access without additional synchronization.
func (q *Queue[T]) Sync(enable bool) {
	q.sync = enable
}

// IsSync returns true if the collection is configured for synchronization,
// meaning that concurrent access is synchronized using locks.
// Returns false if synchronization is disabled,
// allowing concurrent access without additional synchronization.
func (q *Queue[T]) IsSync() bool {
	return q.sync
}

// Represents an enqueue operation.
type EnqOp[T any] struct {
	q *Queue[T]
	v []T
}

func (op *EnqOp[T]) Head() {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	op.q.v = append(op.v, op.q.v...)
}

func (op *EnqOp[T]) Tail() {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	op.q.v = append(op.q.v, op.v...)
}

type DeqOp[T any] struct {
	q *Queue[T]
}

func (op *DeqOp[T]) Head() T {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	result := op.q.v[0]
	op.q.v = op.q.v[1:op.q.Len()]
	return result
}

func (op *DeqOp[T]) TryHead() (T, error) {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	if len(op.q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := op.q.v[0]
	op.q.v = op.q.v[1:op.q.Len()]
	return result, nil
}

func (op *DeqOp[T]) Tail() T {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	result := op.q.v[len(op.q.v)-1]
	op.q.v = op.q.v[0 : op.q.Len()-1]
	return result
}

func (op *DeqOp[T]) TryTail() (T, error) {
	if op.q.IsSync() {
		defer op.q.mu.Lock()
		op.q.mu.Unlock()
	}

	if len(op.q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := op.q.v[len(op.q.v)-1]
	op.q.v = op.q.v[0 : op.q.Len()-1]
	return result, nil
}

type PeekOp[T any] struct {
	q *Queue[T]
}

func (op *PeekOp[T]) Head() T {
	if op.q.IsSync() {
		defer op.q.mu.RLock()
		op.q.mu.RUnlock()
	}

	result := op.q.v[0]
	return result
}

func (op *PeekOp[T]) TryHead() (T, error) {
	if op.q.IsSync() {
		defer op.q.mu.RLock()
		op.q.mu.RUnlock()
	}

	if len(op.q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := op.q.v[0]
	return result, nil
}

func (op *PeekOp[T]) Tail() T {
	if op.q.IsSync() {
		defer op.q.mu.RLock()
		op.q.mu.RUnlock()
	}

	result := op.q.v[len(op.q.v)-1]
	return result
}

func (op *PeekOp[T]) TryTail() (T, error) {
	if op.q.IsSync() {
		defer op.q.mu.RLock()
		op.q.mu.RUnlock()
	}

	if len(op.q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := op.q.v[len(op.q.v)-1]
	return result, nil
}
