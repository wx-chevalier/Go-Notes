
package main

import (
	prettyslice "github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true

	prettyslice.Show("make([]int, 3)", make([]int, 3))
	prettyslice.Show("make([]int, 3, 5)", make([]int, 3, 5))

	s := make([]int, 0, 5)
	prettyslice.Show("make([]int, 0, 5)", s)

	s = append(s, 42)
	prettyslice.Show("s = append(s, 42)", s)
}
