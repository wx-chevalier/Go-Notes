
package main

import "fmt"

func main() {
	var n int

	// ALTERNATIVES:
	// n = n + 1
	// n += 1

	// BETTER:
	n++

	fmt.Println(n)
}
