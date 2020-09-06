
package main

import "fmt"

type product struct {
	title string
	price money
}

func (p *product) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
