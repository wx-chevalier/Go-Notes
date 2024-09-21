
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get and split the PATH environment variable

	// SplitList function automatically finds the
	// separator for the path env variable
	words := filepath.SplitList(os.Getenv("PATH"))

	// Alternative way, but above one is better:
	// words := strings.Split(
	// 	os.Getenv("PATH"),
	// 	string(os.PathListSeparator))

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			q, w = strings.ToLower(q), strings.ToLower(w)

			if strings.Contains(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}
