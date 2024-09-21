
package main

import (
	"fmt"
	"os"
)

func main() {
	// WARNING: This program will error
	//          if you don't pass your name and lastname

	name, lastname := os.Args[1], os.Args[2]

	msg := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, name, lastname)
}
