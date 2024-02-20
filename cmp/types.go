package cmp

type Equatable[T any] interface {
	Equals(other T) bool
}

type Lesser[T any] interface {
	Less(other T) bool
}

type Comparable[T any] interface {
	Compare(other T) int
}
