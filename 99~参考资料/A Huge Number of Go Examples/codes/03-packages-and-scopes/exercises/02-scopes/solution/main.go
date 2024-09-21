
package main

import "fmt"

func main() {
	// as you can see, I don't need to import a package
	// and I can call `hello` function here.
	//
	// this is because, package-scoped names
	//   are shared in the same package
	hello()

	// but here, i can't access the fmt package without
	// importing it.
	//
	// this is because, it's in the printer.go's file scope.
	// it imports it.

	// this main func can also call bye function here
	// bye()
}

// printer.go can call this function
//
// this is because, bye function is in the package-scope
// of the main package now.
//
// main func can also call this.
func bye() {
	fmt.Println("bye bye")
}
