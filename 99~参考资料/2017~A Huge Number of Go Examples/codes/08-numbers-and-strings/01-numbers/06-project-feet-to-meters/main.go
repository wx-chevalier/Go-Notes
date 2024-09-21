
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	// feet is a float64 now
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%f feet is %f meters.\n", feet, meters)

	// pretty print it:

	// fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
