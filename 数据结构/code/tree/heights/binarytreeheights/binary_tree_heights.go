package binarytreeheights

import (
	"github.com/Dev-Snippets/algorithm-go-snippets/datastructures/trees"
)

// Height returns a tree's height
func Height(node *trees.BinaryNode) int {
	if node == nil {
		return 0
	}

	left := Height(node.Left)
	right := Height(node.Right)

	if left > right {
		return 1 + left
	}
	return 1 + right
}
