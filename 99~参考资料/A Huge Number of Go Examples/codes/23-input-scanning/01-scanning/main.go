
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Simulate an error
	// os.Stdin.Close()

	// Create a new scanner that scans from the standard-input
	in := bufio.NewScanner(os.Stdin)

	// Stores the total number of lines in the input
	var lines int

	// Scan the input line by line
	for in.Scan() {
		lines++

		// Get the current line from the scanner's buffer
		// fmt.Println("Scanned Text :", in.Text())
		// fmt.Println("Scanned Bytes:", in.Bytes())
		in.Text()
	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
