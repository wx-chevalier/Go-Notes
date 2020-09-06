
package main

import "fmt"

func main() {
	// When argument to len is a constant, len can be used
	// while initializing a constant
	//
	// Here, "Hello" is a constant.

	const max int = len("Hello")

	fmt.Println(max)
}
