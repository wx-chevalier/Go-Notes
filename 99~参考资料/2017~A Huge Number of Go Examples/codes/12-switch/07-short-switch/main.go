
package main

import "fmt"

func main() {
	// i := 10

	// true is in a comment now
	// you can delete that part if you want
	// it's by default true anyway.
	switch i := 10; /* true */ {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
