
package main

import "fmt"

type money float64

// String() returns a string representation of money.
// money is an fmt.Stringer.
func (m money) String() string {
	return fmt.Sprintf("$%.2f", m)
}
