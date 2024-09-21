
package main

import "fmt"

func main() {
	rates := [3]float64{
		// index 0 empty
		// index 1 empty
		2: 1.5, // index: 2
	}

	fmt.Println(rates)

	// above array literal equals to this:
	//
	// rates := [3]float64{
	// 	0.,
	// 	0.,
	// 	1.5,
	// }
}
