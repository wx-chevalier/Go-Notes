
package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// You can access functions from other files
	// which are in the same package

	// Here, `main()` can access `bye()` and `hey()`

	// It's because bye.go, hey.go and main.go
	//   are in the main package.

	bye()
	hey()
}
