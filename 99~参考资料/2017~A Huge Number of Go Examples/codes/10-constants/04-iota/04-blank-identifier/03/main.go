
package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		_
		MST // CORRECT: -7
		PST // CORRECT: -8
	)

	fmt.Println(EST, MST, PST)
}
