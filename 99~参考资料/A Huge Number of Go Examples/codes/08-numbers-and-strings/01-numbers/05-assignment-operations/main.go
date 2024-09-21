
package main

import (
	"fmt"
)

func main() {
	width, height := 5., 12.

	// calculates the area of a rectangle
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // decreases area by 10
	area = area + 10 // increases area by 10
	area = area * 2  // doubles the area
	area = area / 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// // ASSIGNMENT OPERATIONS
	area -= 10 // decreases area by 10
	area += 10 // increases area by 10
	area *= 2  // doubles the area
	area /= 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// finds the remainder of area variable
	// since: area is float, this won't work:
	// area %= 7

	// this works
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}
