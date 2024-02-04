package sliceutil

import "github.com/sgago/algomon/comp"

func Swap[T any](col []T, idxA, idxB int) {
	col[idxB], col[idxA] = col[idxA], col[idxB]
}

func Min[T comp.Types](vals ...T) T {
	min := vals[0]

	for _, val := range vals[1:] {
		if min > val {
			min = val
		}
	}

	return min
}

func Max[T comp.Types](vals ...T) T {
	max := vals[0]

	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}

	return max
}
