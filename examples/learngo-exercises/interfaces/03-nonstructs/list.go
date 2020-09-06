
package main

import "fmt"

// + you can attach methods to non-struct types.
// + rule: you need to declare a new type in the same package.
type list []*game

func (l list) print() {
	// `list` acts like a `[]game`
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}
