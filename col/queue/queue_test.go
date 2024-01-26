package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int](5, 1, 2, 3)

	q.Enq(4, 5, 6).Head()
	q.Enq(7, 8, 9).Tail()
	q.Deq().Head()
}
