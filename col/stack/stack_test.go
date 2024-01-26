package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s2 := New[int](10)
	fmt.Println(s2.String())

	s := New[int](5, 5, 4, 3, 2, 1)

	str := s.String()
	fmt.Println(str)
}
