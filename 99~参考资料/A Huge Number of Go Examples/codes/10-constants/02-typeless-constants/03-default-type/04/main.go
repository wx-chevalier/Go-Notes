
package main

import "fmt"

func main() {
	i := 42 + 3.14 // OK: 42 and 3.14 are typeless

	// above line equals to this:
	// i := float64(42 + 3.14)

	fmt.Printf("i is %T\n", i)
}
