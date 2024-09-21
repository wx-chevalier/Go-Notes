
package main

// file scope
import "fmt"

// package scope
const ok = true

// package scope
func main() { // block scope starts

	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)

} // block scope ends
