
package main

import "fmt"

func main() {
	// Below solutions are correct:
	x := 5. / 2
	// x := 5 / 2.
	// x := float64(5) / 2
	// x := 5 / float64(2)

	fmt.Println(x)
}
