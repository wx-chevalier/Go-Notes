
package main

import "fmt"

func main() {
	// uint16 max value is 65535
	big := uint16(65535)

	// uint8 destroys its value
	// to its own max value which is 255
	//
	// 65535 - 255 is lost.
	small := uint8(big)

	// fmt.Printf("%b %d\n", big, big)
	// fmt.Printf("%b %[1]d\n", big)

	fmt.Printf("%016b %[1]d\n", big)
	fmt.Printf("%016b %[1]d\n", small)
}
