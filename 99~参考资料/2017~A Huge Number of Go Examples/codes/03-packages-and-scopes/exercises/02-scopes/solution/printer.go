
package main

import "fmt"

func hello() {
	// only this file can access the imported fmt package
	// when others also do so, they can also access
	//   their own `fmt` "name"

	fmt.Println("hi! this is inanc!")
	bye()
}
