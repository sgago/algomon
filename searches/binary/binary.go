package binary

import (
	"github.com/sgago/algomon/comp"
	"github.com/sgago/algomon/errs"
)

// SearchFunc searches a sorted or semi-sorted array for a target value or boundary condition, respectively.
//
// There are two types of binary searches we can perform.
//
//  1. The input array contains the exact value we are searching for.
//     This binary search is straightforward, if you find the target, return immediately.
//
//  2. We're looking for some boundary condition: index of some number smaller than a given value, the first occurrence, etc.
//
//     Time complexity is O(logN) since we cut the solution space in half each time.
func SearchFunc[T any](sorted []T, compare func(idx int, val T) int) (int, T) {
	mid, left, right := 0, 0, len(sorted)-1

	for left <= right {
		mid = left + (right-left)/2

		comp := compare(mid, sorted[mid])

		if comp == 0 {
			return mid, sorted[mid] // Found it!
		} else if comp < 0 {
			// val is < 0, so pull left side in + 1
			left = mid + 1
		} else {
			// val is > 0, so pull right side in - 1
			right = mid - 1
		}
	}

	if left < 0 || left >= len(sorted) {
		return -1, *new(T)
	}

	return left, sorted[left]
}

func TrySearchFunc[T any](sorted []T, predicate func(idx int, val T) int) (int, T, error) {
	idx, val := SearchFunc(sorted, predicate)

	if idx == -1 {
		return -1, val, errs.NotFound
	}

	return idx, val, nil
}

// Search searches a sorted array for a target value.
//
//	Time complexity is O(logN) since we cut the solution space in half each time.
func Search[T comp.Types](sorted []T, target T) int {
	idx, _ := SearchFunc(sorted, func(idx int, val T) int {
		if val == target {
			return 0
		} else if val < target {
			return -1
		}

		return 1
	})

	return idx
}

func TrySearch[T comp.Types](arr []T, target T) (int, error) {
	idx := Search(arr, target)

	if idx == -1 {
		return idx, errs.NotFound
	}

	return idx, nil
}
