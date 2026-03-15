
package main

import "fmt"

func main() {
	// I'm using multiple declarations instead of singular
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
}
