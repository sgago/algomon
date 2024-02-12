package stack

import (
	"fmt"

	"github.com/sgago/algomon/col/conc"
	"github.com/sgago/algomon/sliceutil"
)

type IdxStack[T any] struct {
	conc.RWLocker

	v   []T
	idx int
}

func NewIdxStack[T any](cap int, vals ...T) *IdxStack[T] {
	s := &IdxStack[T]{
		v:   make([]T, cap),
		idx: -1,
	}

	s.Push(vals...)

	return s
}

func (s *IdxStack[T]) Push(vals ...T) {
	if len(vals) == 0 {
		return
	}

	defer s.RWLocker.Write()()

	newIdx := s.idx + len(vals)

	if newIdx >= len(s.v) {
		s.v = append(s.v, make([]T, len(s.v)+len(vals))...)
	}

	sliceutil.Reverse(vals)

	copy(s.v[s.idx+1:], vals)

	s.idx = newIdx
}

func (s *IdxStack[T]) Pop() T {
	defer s.RWLocker.Write()()
	result := s.v[s.idx]
	s.idx--
	return result
}

func (s *IdxStack[T]) Slice() []T {
	defer s.RWLocker.Write()()
	return s.v[:s.idx+1]
}

func (s *IdxStack[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

func (s *IdxStack[T]) Len() int {
	defer s.RWLocker.Read()()
	return s.idx + 1
}

func (s *IdxStack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *IdxStack[T]) Resize(cap int) {
	defer s.RWLocker.Write()()
	s.v = s.v[:max(len(s.v), cap)]
}
