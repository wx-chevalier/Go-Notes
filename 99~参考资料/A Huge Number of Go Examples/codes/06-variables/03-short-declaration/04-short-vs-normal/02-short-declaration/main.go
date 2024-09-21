
package main

// short declaration use cases

func main() {

	// -----------------------------------------------------
	// if you know the initial value
	// -----------------------------------------------------

	// DON'T DO THIS:
	// var width, height = 100, 50

	// DO THIS (concise):
	// width, height := 100, 50

	// -----------------------------------------------------
	// redeclaration
	// -----------------------------------------------------

	// DON'T DO THIS:
	// width = 50
	// color := red

	// DO THIS (concise):
	// width, color := 50, "red"
}
