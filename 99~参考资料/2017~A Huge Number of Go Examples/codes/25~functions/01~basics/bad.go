
package main

import "fmt"

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("showN       : N is %d\n", N)
}

func incrN() {
	N++
}

// you cannot declare a function within the same package with the same name
// func incrN() {
// }
