
package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	//
	// each int element is 4 bytes (on 64-bit)
	//
	// let's say the ages point to 1000th.
	// ages[1:] will point to 1004th
	// ages[2:] will point to 1008th and so on.
	//
	// they all will be looking at the same
	// backing array.
	//

	ages := []int{35, 15, 25}
	red, green := ages[0:1], ages[1:3]

	s.Show("ages", ages)
	s.Show("red", red)
	s.Show("green", green)

	fmt.Println(red[0])
	// fmt.Println(red[1]) // error
	// fmt.Println(red[2]) // error

	{
		var ages []int
		s.Show("nil slice", ages)

		// or just:
		s.Show("nil slice", []int(nil))
	}
}
