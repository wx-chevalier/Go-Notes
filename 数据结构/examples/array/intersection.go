package main

import (
	"fmt"
)

/**
* 给定两个数组，编写一个函数来计算它们的交集。
* 输入: nums1 = [1,2,2,1], nums2 = [2,2]
* 输出: [2,2]
 */
func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	resp := []int{}

	for _, v := range nums1 {
		m0[v] = 1
	}

	for _, v := range nums2 {

		if m0[v] > 0 {
			m0[v]--
			resp = append(resp, v)
		}
	}

	return resp
}

func main() {
	resp := intersect([]int{1, 2, 2, 3}, []int{1, 2, 4, 5, 6})

	fmt.Print(resp)
}
