
package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

func main() {
	msg := os.Args[1]

	s := msg + strings.Repeat("!", len(msg))

	fmt.Println(s)
}
