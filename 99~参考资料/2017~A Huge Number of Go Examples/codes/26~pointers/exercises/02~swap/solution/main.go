
package main

import "fmt"

func main() {
	a, b := 3.14, 6.28
	swap(&a, &b)
	fmt.Printf("a : %g — b : %g\n", a, b)

	pa, pb := &a, &b
	pa, pb = swapAddr(pa, pb)
	fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	fmt.Printf("pa: %g — pb: %g\n", *pa, *pb)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func swapAddr(a, b *float64) (*float64, *float64) {
	return b, a
}
