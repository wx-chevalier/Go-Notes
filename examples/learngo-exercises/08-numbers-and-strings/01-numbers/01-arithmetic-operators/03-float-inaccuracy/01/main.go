
package main

import "fmt"

func main() {
	ratio := 1.0 / 10.0

	// after 10 operations
	// the inaccuracy is clear
	//
	// BTW, don't mind about this loop syntax for now
	// I'm going to explain it afterwards
	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f", ratio)
}
