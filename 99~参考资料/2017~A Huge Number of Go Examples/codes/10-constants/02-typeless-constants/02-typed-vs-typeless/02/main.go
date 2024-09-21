
package main

import "fmt"

func main() {
	const min int = 42

	var f float64

	// ERROR: Type Mismatch
	// f = min // NOT OK

	fmt.Println(f)
}
