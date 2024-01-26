package casstack

// Reference: https://www.geeksforgeeks.org/lock-free-stack-using-java/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestStackC(t *testing.T) {
	s := CasStack[int]{}

	s.Push(123)
	result := s.Pop()

	fmt.Printf("%v", result)
}

func TestStackCPopConcurrent(t *testing.T) {
	const stackSize = 1000

	s := CasStack[int]{}

	// Populate the stack with numbers from 1 to stackSize
	for i := 0; i < stackSize; i++ {
		s.Push(i)
	}

	var wg sync.WaitGroup

	m := make([]*atomic.Int32, stackSize)
	for i := 0; i < stackSize; i++ {
		m[i] = &atomic.Int32{}
	}

	for i := 0; i < stackSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			popped := s.Pop()
			m[popped].Add(1)
		}()
	}

	wg.Wait()

	// Verify that all values from [0 to stackSize) were popped exactly once
	for i := 0; i < stackSize; i++ {
		count := m[i].Load()

		if count > 1 {
			t.Errorf("Value %v was popped %d times", i, count)
		}
	}
}
