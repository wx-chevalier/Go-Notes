
package main

import "fmt"

func main() {
	// constants repeat the previous type and expression
	const (
		min int = 1
		max     // int = 1
	)

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}
