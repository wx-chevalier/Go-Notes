
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Itoa doesn't return any errors
	// So, you don't have to handle the errors for it

	s := strconv.Itoa(42)

	fmt.Println(s)
}
