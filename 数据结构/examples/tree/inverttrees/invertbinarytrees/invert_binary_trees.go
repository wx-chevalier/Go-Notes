package invertbinarytrees

import "github.com/Dev-Snippets/algorithm-go-snippets/datastructures/trees"

// InvertTree inverts a binary tree
func InvertTree(node *trees.BinaryNode) *trees.BinaryNode {
	if node == nil {
		return nil
	}

	node.Left, node.Right = InvertTree(node.Right), InvertTree(node.Left)
	return node
}
