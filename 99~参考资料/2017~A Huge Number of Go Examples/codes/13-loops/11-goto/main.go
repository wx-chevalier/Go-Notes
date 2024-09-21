
package main

import (
	"fmt"
)

func main() {
	// cannot step over any variable declarations
	// ERROR: "i" variable is declared after the jump
	//
	// goto loop

	var i int

loop:
	if i < 3 {
		fmt.Println("looping")
		i++
		goto loop
	}
	fmt.Println("done")
}
