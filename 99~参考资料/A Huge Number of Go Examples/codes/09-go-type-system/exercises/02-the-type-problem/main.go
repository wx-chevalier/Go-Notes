
package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: The Type Problem
//
//  Solve the data type problem in the program.
//
// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true
// ---------------------------------------------------------

func main() {
	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.
	var (
		width  uint8
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// UNCOMMENT THIS:
	// fmt.Println("are they equal?", width == height)
}
