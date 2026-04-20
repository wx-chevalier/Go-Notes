
package main

import "fmt"

func main() {
	// min and pi are typeless constants
	const min = 1 + 1
	const pi = 3.14 * min

	fmt.Println(min, pi)
}
