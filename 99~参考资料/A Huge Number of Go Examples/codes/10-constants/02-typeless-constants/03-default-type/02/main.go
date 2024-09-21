
package main

import "fmt"

func main() {
	const min = 1

	max := 5 + min
	// above line equals to this:
	// max := int(5) + int(min)

	fmt.Printf("Type of max: %T\n", max)
}
