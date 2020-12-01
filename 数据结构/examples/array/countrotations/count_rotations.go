package countrotations

import "github.com/Dev-Snippets/algorithm-go-snippets/slices/rotations/minrotations"

// Count counts the number of rotations
func Count(array []int) int {
	minimumIndex := minrotations.Min(array)
	if minimumIndex == 0 {
		return 0
	}
	return len(array) - minimumIndex
}
