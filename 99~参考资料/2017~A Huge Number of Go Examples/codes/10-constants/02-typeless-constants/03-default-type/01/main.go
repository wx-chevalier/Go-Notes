
package main

import "fmt"

func main() {
	const min int32 = 1

	max := 5 + min
	// above line equals to this:
	// max := int32(5) + min

	fmt.Printf("Type of max: %T\n", max)
}
