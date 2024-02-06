package stack

import (
	"sync"
	"sync/atomic"
	"testing"

	normal "github.com/sgago/algomon/col/stack"
)

const num = 5000

var nonSyncStack *normal.Stack[int]
var syncStack *normal.Stack[int]
var casStack *Stack[int]
var nonSyncIdxStack *normal.IdxStack[int]
var syncIdxStack *normal.IdxStack[int]

func init() {
	nonSyncStack = normal.New[int](num)
	nonSyncIdxStack = normal.NewIdxStack[int](num)

	syncStack = normal.New[int](num)
	syncStack.Synchronize(true)
	syncIdxStack = normal.NewIdxStack[int](num)
	syncIdxStack.Synchronize(true)

	casStack = New[int]()
}

func BenchmarkNonSyncStackPushPop(b *testing.B) {
	s := nonSyncStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkNonSyncStackStringify(b *testing.B) {
	s := nonSyncStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkNonSyncIdxStackPushPop(b *testing.B) {
	s := nonSyncIdxStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkNonSyncIdxStackStringify(b *testing.B) {
	s := nonSyncIdxStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkSyncIdxStackPushPop(b *testing.B) {
	s := syncIdxStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkSyncIdxStackStringify(b *testing.B) {
	s := syncIdxStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkSyncIdxStackConcurrent(b *testing.B) {
	s := syncIdxStack

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < num; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, num)
	for i := 0; i < num; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()
}

func BenchmarkSyncStackPushPop(b *testing.B) {
	s := syncStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkSyncStackStringify(b *testing.B) {
	s := syncStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkSyncStackConcurrent(b *testing.B) {
	s := syncStack

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < num; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, num)
	for i := 0; i < num; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()
}

func BenchmarkCasStackPushPop(b *testing.B) {
	s := casStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	for !s.Empty() {
		s.Pop()
	}
}

func BenchmarkCasStackStringify(b *testing.B) {
	s := casStack

	for i := 0; i < num; i++ {
		s.Push(i)
	}

	_ = s.String()
}

func BenchmarkCasStackConcurrent(b *testing.B) {
	s := casStack

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < num; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, num)
	for i := 0; i < num; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()
}
