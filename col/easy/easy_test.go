package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {
	s := Stack{}

	s.Push(12)
	s.Push(34)
	s.Push(56)

	assert.Equal(t, 56, s.Pop())
	assert.Equal(t, 34, s.Pop())
	assert.Equal(t, 12, s.Pop())
}

func TestEnqueueDequeue(t *testing.T) {
	q := Queue{}

	q.Enq(12)
	q.Enq(34)
	q.Enq(56)

	assert.Equal(t, 12, q.Deq())
	assert.Equal(t, 34, q.Deq())
	assert.Equal(t, 56, q.Deq())
}
