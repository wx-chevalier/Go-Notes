
package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
//
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

func main() {
	msg := `
	
	The weather looks good.
I should go and play.



	`

	fmt.Println(msg)
}
