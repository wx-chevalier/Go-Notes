
package main

import "fmt"

func main() {
	n, m := 1, 5

	fmt.Println(2 + 1*m/n)
	fmt.Println(2 + ((1 * m) / n)) // same as above

	// let's change the precedence using parentheses
	fmt.Println(((2 + 1) * m) / n)
}
