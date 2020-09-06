
package main

import "fmt"

func main() {
	// OPTION #1 (option #2 is better)
	// var safe bool = true

	// OPTION #2 (OK)
	// var safe = true

	// OPTION #3 - SHORT DECLARATION (BEST)
	//
	// You don't even need to type the `var` keyword
	//
	// Short declaration equals to:
	//   var safe bool = true
	//   var safe = true
	//
	// Go gets (infers) the type from the initializer value
	//
	// true's default type is bool
	// so, the type of the safe variable becomes a bool
	safe := true

	fmt.Println(safe)
}
