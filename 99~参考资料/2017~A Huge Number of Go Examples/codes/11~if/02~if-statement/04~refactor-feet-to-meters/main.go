
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))

		// ALTERNATIVE:
		// fmt.Println("Please tell me a value in feet")

		return
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
