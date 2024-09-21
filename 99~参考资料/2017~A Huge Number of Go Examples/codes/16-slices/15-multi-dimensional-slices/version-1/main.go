
package main

import (
	"fmt"
)

func main() {
	// 1st day: $200, $100
	// 2nd day: $500
	// 3rd day: $50, $25, and $75

	spendings := [][]int{
		{200, 100},   // 1st day
		{500},        // 2nd day
		{50, 25, 75}, // 3rd day
	}

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}
