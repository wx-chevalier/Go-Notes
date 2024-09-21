
package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}
