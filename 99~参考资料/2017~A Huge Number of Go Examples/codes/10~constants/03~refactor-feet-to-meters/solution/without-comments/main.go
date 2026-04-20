
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// I copied the same program here but without comments.
// So, you can easily read it.

func main() {
	const (
		feetInMeters float64 = 0.3048
		feetInYards          = feetInMeters / 0.9144
	)

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
