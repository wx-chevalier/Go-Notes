package duplicates

import "github.com/Dev-Snippets/algorithm-go-snippets/datastructures/sets/hashsets"

// ContainsDuplicates checks if the list contains duplicates
func ContainsDuplicates(values ...interface{}) bool {
	set := hashsets.New()
	for _, value := range values {
		if set.Contains(value) {
			return true
		}
		set.Add(value)
	}
	return false
}
