package dsu

type node[T any] struct {
	val    any
	parent *node[T]
}

type Dsu[T any] struct {
	nodes map[int]*node[T]
}
