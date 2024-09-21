
package main

import "fmt"

func main() {
	{
		// its length is part of its type
		var nums [5]int
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		// its length is not part of its type
		var nums []int
		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

		// won't work: the slice is nil.
		// fmt.Printf("nums[0]: %d\n", nums[0])
		// fmt.Printf("nums[1]: %d\n", nums[1])
	}
}
