
package main

import "fmt"

func main() {
	speed := 100 // int
	force := 2.5 // float64

	// ERROR: invalid op
	// speed = speed * force

	// conversion can be a destructive operation
	// `force` loses its fractional part...

	speed = speed * int(force)

	fmt.Println(speed)
}
