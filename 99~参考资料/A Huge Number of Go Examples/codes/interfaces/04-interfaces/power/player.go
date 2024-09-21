
package main

import (
	"fmt"
)

// Player can be powered
type Player struct{}

// Draw power to a Player
func (Player) Draw(power int) {
	fmt.Printf("Player is drawing %dkW of electrical power.\n", power)
}
