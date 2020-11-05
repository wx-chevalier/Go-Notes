package groupanagrams

import (
	"github.com/Dev-Snippets/algorithm-go-snippets/datastructures/maps/hashmultimaps"
	"github.com/Dev-Snippets/algorithm-go-snippets/strings/sorts"
)

// GroupAnagrams groups anagrams
func GroupAnagrams(list []string) *hashmultimaps.HashMultiMap {
	multimap := hashmultimaps.New()
	for _, value := range list {
		sortedValue := sorts.Sort(value)
		multimap.Put(sortedValue, value)
	}
	return multimap
}
