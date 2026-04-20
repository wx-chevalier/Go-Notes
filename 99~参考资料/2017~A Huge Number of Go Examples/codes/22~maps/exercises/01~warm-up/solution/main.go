
package main

import "fmt"

func main() {
	var (
		phones map[string]string
		// Key        : Last name
		// Element    : Phone number

		// Key        : Product ID
		// Element    : Available / Unavailable
		products map[int]bool

		multiPhones map[string][]string
		// Key        : Last name
		// Element    : Phone numbers

		basket map[int]map[int]int
		// Key        : Customer ID
		// Element Key:
		//   Key: Product ID Element: Quantity
	)

	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", multiPhones)
	fmt.Printf("basket     : %#v\n", basket)
}
