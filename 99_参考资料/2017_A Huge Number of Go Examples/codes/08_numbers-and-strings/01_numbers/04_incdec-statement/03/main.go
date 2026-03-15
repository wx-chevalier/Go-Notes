
package main

import "fmt"

// incdec is a statement

func main() {
	var counter int

	// following "statements" are correct:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	fmt.Printf("There are %d line(s) in the file\n",
		counter)

	// the following "expressions" are incorrect:

	// counter = 5+counter--
	// counter = ++counter + counter--
}
