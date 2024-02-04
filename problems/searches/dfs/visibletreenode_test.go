/*
Visible Tree Node
https://algo.monster/problems/visible_tree_node

In a binary tree, a node is labeled as "visible" if,
on the path from the root to that node,
there isn't any node with a value higher than this node's value.

Input:
    5
   / \
  4   6
 / \
3   8

Output:
There are 3 visible nodes: 5, 6, 8.

The root is always "visible" since there are no other nodes between the root and itself.
Given a binary tree, count the number of "visible" nodes.
*/

package dfs

import (
	"math"
	"testing"

	"github.com/sgago/algomon/col/node/binary"
	"github.com/sgago/algomon/col/tree/bst"
	"github.com/sgago/algomon/sliceutil"
)

func TestVisibleNodes(t *testing.T) {
	tree := bst.New(2, 1, 3)

	visibleNode(tree.Root, math.MinInt)
}

func visibleNode(n *binary.Node[int], max int) int {
	result := 0

	if n.Leaf() {
		if n.Val > max {
			return 1
		}

		return 0
	}

	if n.Val > max {
		result++
	}

	max = sliceutil.Max(max, n.Val)

	if n.Left != nil {
		result += visibleNode(n.Left, max)
	}

	if n.Right != nil {
		result += visibleNode(n.Right, max)
	}

	return result
}
