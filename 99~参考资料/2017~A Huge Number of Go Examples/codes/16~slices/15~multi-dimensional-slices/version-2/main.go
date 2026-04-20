
package main

import (
	"fmt"
)

func main() {
	spendings := make([][]int, 0, 5)

	spendings = append(spendings, []int{200, 100})
	spendings = append(spendings, []int{25, 10, 45, 60})
	spendings = append(spendings, []int{5, 15, 35})
	spendings = append(spendings, []int{95, 10}, []int{50, 25})

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}
