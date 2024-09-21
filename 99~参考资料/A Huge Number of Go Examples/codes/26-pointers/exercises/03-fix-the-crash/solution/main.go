
package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	c := &computer{} // init with a value (before: c was nil)
	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand) // print the pointed value
}

func change(c *computer, brand string) {
	c.brand = &brand // set the brand's address
}
