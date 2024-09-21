
package main

import (
	"fmt"
)

// Blender can be powered
type Blender struct{}

// Draw power to a Blender
func (Blender) Draw(power int) {
	fmt.Printf("Blender is drawing %dkW of electrical power.\n", power)
}
