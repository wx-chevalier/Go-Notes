
package main

import (
	"fmt"
)

// Kettle can be powered
type Kettle struct{}

// Draw power to a Kettle
func (Kettle) Draw(power int) {
	fmt.Printf("Kettle is drawing %dkW of electrical power.\n", power)
}
