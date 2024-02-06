package steve

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNSort(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	actual := nsort(1, arr)
	assert.True(t, slices.IsSorted(actual))
}

func TestQuick(t *testing.T) {
	arr := []int{5, 3, 1, 4, 2}
	actual := quick(arr)
	assert.True(t, slices.IsSorted(actual))
}

func TestNSortGaps(t *testing.T) {
	// 1 2 3 4 5 <- perfect
	// 1 2 3 <-> 7 8 9 <- gap of 4
	// Min: 1, Max: 9, N: 6 (we can compute the expected sum for a given N)
	// Sum: 30, expected sum with min is 15
	// We can compute the arithmetic sum of the series with (max + min)*N/2 or some such.
	// I wonder if we can use clever bit shifting to pile up buckets without gaps...
	// We an also sum the number of ordered vs unordered elements and capatilize on this.
	arr := []int{3, 9, 2, 8, 1, 7}
	actual := nsort(1, arr)
	assert.True(t, slices.IsSorted(actual))
}
