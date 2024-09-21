
package main

import (
	"fmt"
	"os"
)

// Since, you didn't learn about the control flow statements yet
// I didn't include an error detection here.
//
// So, if you don't pass a name and a lastname,
// this program will fail.
// This is intentional.

func main() {
	var name string

	// assign a new value to the string variable below
	name = os.Args[1]
	fmt.Println("Hello great", name, "!")

	// changes the name, declares the age with 85
	name, age := "gandalf", 2019

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
