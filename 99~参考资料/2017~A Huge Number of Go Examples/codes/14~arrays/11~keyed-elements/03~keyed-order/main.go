
package main

import "fmt"

func main() {
	rates := [3]float64{
		1: 2.5, // index: 1
		0: 0.5, // index: 0
		2: 1.5, // index: 2
	}

	fmt.Println(rates)

	// above array literal equals to this:
	//
	// rates := [3]float64{
	// 	0.5,
	// 	2.5,
	// 	1.5,
	// }
}
