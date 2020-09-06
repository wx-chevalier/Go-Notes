
package main

import "fmt"

func main() {
	city := "Paris"

	switch city {
	case "Paris":
		fmt.Println("France")
	}

	// SIMILAR TO IF
	// ------------------------------------

	// switch statement is converted to an if statement
	// automatically behind the scenes
	//
	// above switch statement is similar to this if

	// if city == "Paris" {
	// 	fmt.Println("France")
	// }
}
