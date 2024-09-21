
package main

import "fmt"

func main() {
	n := 10

	// ALTERNATIVES:
	// n = n - 1
	// n -= 1

	// BETTER:
	n--

	fmt.Println(n)
}
