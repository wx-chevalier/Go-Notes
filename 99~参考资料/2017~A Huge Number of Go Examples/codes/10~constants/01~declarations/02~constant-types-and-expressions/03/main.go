
package main

import "fmt"

func main() {
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1" + "-beta"
	const debug = !true
	const A rune = 'A' + 1

	fmt.Println(min, pi, version, debug, A)
}
