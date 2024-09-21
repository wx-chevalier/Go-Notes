
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {

		// only a is available here
		fmt.Println("Give me a number.")

		// no need to return on error
		// return

	} else if n, err := strconv.Atoi(a[1]); err != nil {

		// a, n and err are available here
		fmt.Printf("Cannot convert %q.\n", a[1])

		// no need to return on error
		// return

	} else {
		// all of the variables (names)
		// are available here
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// a, n and err are not available here
	// they belong to the if statement

	// TRY:
	// fmt.Println(a, n, err)
}
