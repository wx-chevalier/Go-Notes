
package main

import (
	"fmt"
)

// Mixer can be powered
type Mixer struct{}

// Draw power to a Mixer
func (Mixer) Draw(power int) {
	fmt.Printf("Mixer is drawing %dkW of electrical power.\n", power)
}
