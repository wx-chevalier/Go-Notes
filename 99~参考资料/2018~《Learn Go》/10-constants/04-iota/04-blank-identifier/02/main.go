// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		MST               // CORRECT: -6
		PST               // CORRECT: -7
	)

	fmt.Println(EST, MST, PST)
}
