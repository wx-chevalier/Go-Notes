
// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	var c *computer
	change(c, "apple")
	fmt.Printf("brand: %s\n", c.brand)
}

func change(c *computer, brand string) {
	(*c.brand) = brand
}
