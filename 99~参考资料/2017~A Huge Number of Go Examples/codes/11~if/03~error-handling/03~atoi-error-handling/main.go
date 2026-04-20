
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]

	// Atoi returns an int and an error value
	// So, you need to handle the errors

	n, err := strconv.Atoi(age)

	// handle the error immediately and quit
	// don't do anything else here

	if err != nil {
		fmt.Println("ERROR:", err)

		// quits/returns from the main function
		// so, the program ends
		return
	}

	// DO NOT DO THIS:
	// else {
	//   happy path
	// }

	// DO THIS:

	// happy path (err is nil)
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}
