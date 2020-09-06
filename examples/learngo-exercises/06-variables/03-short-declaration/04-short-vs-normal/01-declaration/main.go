
package main

// normal declaration use cases

// -----------------------------------------------------
// when you need a package scoped variable
// -----------------------------------------------------

// version := 0 // YOU CAN'T
var version int

func main() {

	// -----------------------------------------------------
	// if you don't know the initial value
	// -----------------------------------------------------

	// DON'T DO THIS:
	// score := 0

	// DO THIS:
	// var score int

	// -----------------------------------------------------
	// group variables for readability
	// -----------------------------------------------------

	// var (
	// 	video    string

	// 	duration int
	// 	current  int
	// )
}
