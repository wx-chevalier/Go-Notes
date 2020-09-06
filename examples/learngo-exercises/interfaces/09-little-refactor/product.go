
package main

import (
	"fmt"
)

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) print() {
	fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), p.released.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
