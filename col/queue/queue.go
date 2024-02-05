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
	return fmt.Sprintf("%v", q.Slice())
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

func (q *Queue[T]) EnqHead(vals ...T) {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	q.v = append(vals, q.v...)
}

func (q *Queue[T]) EnqTail(vals ...T) {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	q.v = append(q.v, vals...)
}

func (q *Queue[T]) DeqHead() T {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	result := q.v[0]
	q.v = q.v[1:q.Len()]
	return result
}

func (q *Queue[T]) TryDeqHead() (T, error) {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[0]
	q.v = q.v[1:q.Len()]
	return result, nil
}

func (q *Queue[T]) DeqTail() T {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	result := q.v[len(q.v)-1]
	q.v = q.v[0 : q.Len()-1]
	return result
}

func (q *Queue[T]) TryDeqTail() (T, error) {
	if q.IsSync() {
		defer q.mu.Lock()
		q.mu.Unlock()
	}

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[len(q.v)-1]
	q.v = q.v[0 : q.Len()-1]
	return result, nil
}

func (q *Queue[T]) PeekHead() T {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	result := q.v[0]
	return result
}

func (q *Queue[T]) TryPeekHead() (T, error) {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[0]
	return result, nil
}

func (q *Queue[T]) PeekTail() T {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	result := q.v[len(q.v)-1]
	return result
}

func (q *Queue[T]) TryPeekTail() (T, error) {
	if q.IsSync() {
		defer q.mu.RLock()
		q.mu.RUnlock()
	}

	if len(q.v) == 0 {
		return *new(T), errs.Empty
	}

	result := q.v[len(q.v)-1]
	return result, nil
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
