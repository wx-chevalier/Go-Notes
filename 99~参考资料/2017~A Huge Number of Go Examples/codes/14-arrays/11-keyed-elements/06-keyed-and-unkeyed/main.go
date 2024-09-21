
package main

import "fmt"

func main() {
	rates := [...]float64{
		// index 1 to 4 empty

		5:   1.5, // index: 5
		2.5, //      index: 6
		0:   0.5, // index: 0
	}

	fmt.Println(rates)

	// above array literal equals to this:
	//
	// rates := [7]float64{
	// 	0.5,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	1.5,
	//  2.5,
	// }
}
