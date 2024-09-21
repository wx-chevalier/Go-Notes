
package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Pow calculates the power of the given number
	fmt.Println(math.Pow10(2)) // 100
	fmt.Println(math.Pow10(3)) // 1000
	fmt.Println(math.Pow10(4)) // 10000

	// ERROR: math.Pow is not a constant
	//        constants cannot use runtime constructs

	// const max int = math.Pow10(2)
}
