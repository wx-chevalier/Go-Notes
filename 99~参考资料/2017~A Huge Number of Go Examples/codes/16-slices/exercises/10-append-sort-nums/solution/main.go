
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few numbers")
		return
	}

	var nums []int
	for _, s := range args {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		nums = append(nums, n)
	}

	sort.Ints(nums)
	fmt.Println(nums)
}
