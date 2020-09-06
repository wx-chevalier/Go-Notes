
package main

import "fmt"

func main() {
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)
}
