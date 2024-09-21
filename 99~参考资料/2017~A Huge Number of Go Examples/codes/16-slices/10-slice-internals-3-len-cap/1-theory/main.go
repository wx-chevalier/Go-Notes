
package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.MaxPerLine = 6
	s.PrintBacking = true

	ages := []int{35, 15, 25}
	s.Show("ages", ages)

	s.Show("ages[0:0]", ages[0:0])

	for i := 1; i < 4; i++ {
		txt := fmt.Sprintf("ages[%d:%d]", 0, i)
		s.Show(txt, ages[0:i])
	}

	s.Show("append", append(ages, 50))
}
