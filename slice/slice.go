package slice

func Swap[T any](col []T, idxA, idxB int) {
	col[idxB], col[idxA] = col[idxA], col[idxB]
}
