
package main

import "fmt"

func main() {
	// byte is an integer number with 8 bits (1 byte)
	var b byte

	// all bits are empty or 0
	// this is the minimum number a byte can represent
	b = 0
	fmt.Printf("%08b = %d\n", b, b)

	// all bits are full or 1
	// this is the maximum number a byte can represent
	b = 255
	fmt.Printf("%08b = %d\n", b, b)
}
