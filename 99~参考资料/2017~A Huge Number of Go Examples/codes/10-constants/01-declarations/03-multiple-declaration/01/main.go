
package main

import "fmt"

func main() {
	// below declaration is the same as this one:

	// const (
	// 	min int = 1
	// 	max int = 1000
	// )

	const min, max int = 1, 1000

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}
