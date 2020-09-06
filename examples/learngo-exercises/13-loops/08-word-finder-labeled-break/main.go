
package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	// labels and other names do not share the same scope
	// var queries string
	// _ = queries

	// this label labels the parent loop below
	// label's scope is the whole function body
	// not only it's block
queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				break queries
			}
		}
	}
}
