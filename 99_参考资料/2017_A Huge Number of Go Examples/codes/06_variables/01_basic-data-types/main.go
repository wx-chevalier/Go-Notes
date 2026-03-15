
package main

import "fmt"

func main() {
	// integer literal
	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)

	// float literal
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // empty - prints just a space
		"hi",

		// unicode
		"nasılsın?",   // "how are you" in turkish
		"hi there 星!", // "hi there star!"
	)
}
