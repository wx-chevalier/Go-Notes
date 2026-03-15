
package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		MST               // INCORRECT: -6
		PST               // INCORRECT: -7
	)

	fmt.Println(EST, MST, PST)
}
