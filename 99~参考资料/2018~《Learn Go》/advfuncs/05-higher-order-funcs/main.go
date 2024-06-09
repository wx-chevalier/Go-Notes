// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

type filterFunc func(int) bool

func main() {
	fmt.Println("••• HIGHER-ORDER FUNCS •••")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	odd := reverse(reverse(isEven))
	fmt.Printf("reversed   : %t\n", odd(8))

	fmt.Printf("> 3        : %d\n", filter(greater(3), nums...))
	fmt.Printf("> 6        : %d\n", filter(greater(6), nums...))
	fmt.Printf("<= 6       : %d\n", filter(lesseq(6), nums...))
	fmt.Printf("<= 6       : %d\n", filter(reverse(greater(6)), nums...))
}

func reverse(f filterFunc) filterFunc {
	return func(n int) bool {
		return !f(n)
	}
}

func greater(min int) filterFunc {
	return func(n int) bool {
		return n > min
	}
}

func lesseq(max int) filterFunc {
	return func(n int) bool {
		return n <= max
	}
}

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, n := range nums {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return
}

func isEven(n int) bool { return n%2 == 0 }
func isOdd(m int) bool  { return m%2 == 1 }
