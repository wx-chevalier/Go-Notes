
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	spendings := fetch()

	for i, daily := range spendings {
		var total int
		for _, spending := range daily {
			total += spending
		}

		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}

	return spendings
}
