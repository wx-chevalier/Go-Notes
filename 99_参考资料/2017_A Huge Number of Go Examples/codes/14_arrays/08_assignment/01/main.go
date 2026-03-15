
package main

import "fmt"

func main() {
	blue := [3]int{6, 9, 3}
	red := blue

	blue[0] = 10

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", red)

	// UNASSIGNABLE:
	// blue := [3]int{6, 9, 3}
	// red := [2]int{3, 5}
	// red = blue
}
