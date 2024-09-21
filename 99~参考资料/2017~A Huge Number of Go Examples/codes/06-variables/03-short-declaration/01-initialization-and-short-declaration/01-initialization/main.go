
package main

import "fmt"

func main() {
	// = is the assignment operator
	// when used within a variable declaration, it
	// initializes the variable to the given value

	// here, Go initializes the safe variable to true

	// OPTION #1 (option #2 is better)
	// var safe bool = true

	// OPTION #2
	var safe = true

	fmt.Println(safe)
}
