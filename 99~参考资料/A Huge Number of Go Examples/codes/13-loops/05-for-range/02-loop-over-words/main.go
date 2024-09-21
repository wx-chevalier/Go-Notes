
package main

import (
	"fmt"
	"strings"
)

func main() {
	// this will split the following string by spaces
	// and then it will put it inside words variable
	// as a slice of strings
	words := strings.Fields(
		"lazy cat jumps again and again and again",
	)
	//    1    2    3    4    5     6    7    8

	// --------------------------------
	// #1st way:
	// --------------------------------
	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	// --------------------------------
	// #2nd way (best):
	// --------------------------------
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}

	// --------------------------------
	// #3rd way (reuse mechanism):
	// --------------------------------
	// var (
	// 	i int
	// 	v string
	// )

	// for i, v = range words {
	// 	fmt.Printf("#%-2d: %q\n", i+1, v)
	// }

	// fmt.Printf("Last value of i and v %d %q\n", i, v)
}
