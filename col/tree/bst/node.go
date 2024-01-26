package bst

// Node of a BST with a value and pointers to child left and right nodes.
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}
