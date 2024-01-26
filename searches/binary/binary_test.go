package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 7
	actual := Search(arr, expected)
	assert.Equal(t, expected, actual)
}

func TestFind_TargetDoesNotExist(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := Search(arr, 10)
	assert.Equal(t, -1, actual)
}

func TestFindFunc(t *testing.T) {
	target := 2
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	predicate := func(i, v int) int {
		if v == target {
			return 0
		} else if v < target {
			return -1
		}

		return 1
	}

	idx, val := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, target)
	assert.Equal(t, val, target)
}

func TestFind_FirstElementNotSmallerThanTwo(t *testing.T) {
	target := 2
	arr := []int{1, 3, 3, 5, 8, 8, 10}

	predicate := func(i, v int) int {
		if v <= target {
			return -1
		}

		return 1
	}

	idx, val := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, 1)
	assert.Equal(t, val, 3)
}

func TestFind_FirstElementNotSmallerThanNine(t *testing.T) {
	target := 9
	arr := []int{1, 3, 3, 5, 8, 8, 10}

	predicate := func(i, v int) int {
		if v <= target {
			return -1
		}

		return 1
	}

	idx, val := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, 6)
	assert.Equal(t, val, 10)
}

func TestFind_FirstElementNotSmallerThanTwelve(t *testing.T) {
	target := 12
	arr := []int{1, 3, 3, 5, 8, 8, 10}

	predicate := func(i, v int) int {
		if v <= target {
			return -1
		}

		return 1
	}

	idx, _ := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, -1)
}

func TestFind_FirstElementGreaterThanTwo(t *testing.T) {
	target := 2
	arr := []int{1, 3, 3, 5, 8, 8, 10}

	predicate := func(_, v int) int {
		if v > target {
			return 1
		}

		return -1
	}

	idx, val := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, 1)
	assert.Equal(t, val, 3)
}

func TestFind_FirstElementGreaterThanZero(t *testing.T) {
	target := 0
	arr := []int{1, 3, 3, 5, 8, 8, 10}

	predicate := func(_, v int) int {
		if v > target {
			return 1
		}

		return -1
	}

	idx, val := SearchFunc[int](arr, predicate)

	assert.Equal(t, idx, 0)
	assert.Equal(t, val, 1)
}
