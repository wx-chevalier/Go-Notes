
package main

import "fmt"

// This program uses a named constant
// instead of a magic value

func main() {
	const meters int = 100

	cm := 100
	m := cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)
}
