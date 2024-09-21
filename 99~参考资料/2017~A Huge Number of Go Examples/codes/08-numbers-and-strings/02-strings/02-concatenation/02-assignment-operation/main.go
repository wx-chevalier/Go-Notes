
package main

import "fmt"

func main() {
	name, last := "carl", "sagan"

	// assignment operation using string concat
	name += " edward"

	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)
}
