
package main

import (
	"fmt"
)

type product struct {
	Title    string    `json:"title"`
	Price    money     `json:"price"`
	Released timestamp `json:"released"`
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}
