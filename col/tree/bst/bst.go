package bst

import "github.com/sgago/algomon/types"

// Tree represents a binary search tree (BST).
type Tree[T any] struct {
	Root *Node[T]
	less func(a, b T) bool
}

func NewTreeFunc[T any](less func(a, b T) bool, vals ...T) *Tree[T] {
	t := &Tree[T]{
		less: less,
	}

	t.Insert(vals...)

	return t
}

func NewTree[T types.PrimComp](vals ...T) *Tree[T] {
	return NewTreeFunc[T](func(a, b T) bool { return a < b }, vals...)
}

// Inserts values into the BST as nodes.
// Note that duplicates are not allowed in BST.
func (t *Tree[T]) Insert(vals ...T) {
	for _, val := range vals {
		new := &Node[T]{
			Value: val,
		}

		t.Root = t.insert(new, t.Root)
	}
}

func (t *Tree[T]) insert(new *Node[T], next *Node[T]) *Node[T] {
	if next == nil {
		return new
	}

	if t.less(new.Value, (*next).Value) {
		next.Left = t.insert(new, next.Left)
	} else {
		next.Right = t.insert(new, next.Right)
	}

	return next
}

func (t *Tree[T]) InOrder() []T {
	return t.inOrder(t.Root, make([]T, 0))
}

func (t *Tree[T]) inOrder(n *Node[T], results []T) []T {
	if n.Left == nil && n.Right == nil {
		return append(results, n.Value)
	}

	if n.Left != nil {
		results = t.inOrder(n.Left, results)
	}

	results = append(results, n.Value)

	if n.Right != nil {
		results = t.inOrder(n.Right, results)
	}

	return results
}

func (t *Tree[T]) PreOrder() []T {
	return t.preOrder(t.Root, make([]T, 0))
}

func (t *Tree[T]) preOrder(n *Node[T], results []T) []T {
	if n.Left == nil && n.Right == nil {
		return append(results, n.Value)
	}

	results = append(results, n.Value)

	if n.Left != nil {
		results = t.preOrder(n.Left, results)
	}

	if n.Right != nil {
		results = t.preOrder(n.Right, results)
	}

	return results
}

func (t *Tree[T]) PostOrder() []T {
	return t.postOrder(t.Root, make([]T, 0))
}

func (t *Tree[T]) postOrder(n *Node[T], results []T) []T {
	if n.Left == nil && n.Right == nil {
		return append(results, n.Value)
	}

	if n.Left != nil {
		results = t.postOrder(n.Left, results)
	}

	if n.Right != nil {
		results = t.postOrder(n.Right, results)
	}

	results = append(results, n.Value)

	return results
}
