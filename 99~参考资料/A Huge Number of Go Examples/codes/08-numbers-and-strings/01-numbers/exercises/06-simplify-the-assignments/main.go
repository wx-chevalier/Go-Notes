
package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Simplify the Assignments
//
//  Simplify the code (refactor)
//
// RESTRICTION
//  Use only the incdec and assignment operations
//
// EXPECTED OUTPUT
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5

	fmt.Println(width)
}
