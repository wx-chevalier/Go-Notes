
package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	evens := []int{2, 4}
	odds := []int{3, 5, 7}

	s.Show("evens [before]", evens)
	s.Show("odds  [before]", odds)

	N := copy(evens, odds)
	fmt.Printf("%d element(s) are copied.\n", N)

	s.Show("evens [after]", evens)
	s.Show("odds  [after]", odds)
}
