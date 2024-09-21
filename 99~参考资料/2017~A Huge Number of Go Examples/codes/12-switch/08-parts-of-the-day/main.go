
package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now() gets the current time
	// and in turn, .Hour() gets the current hour from it

	// h := time.Now().Hour()
	// fmt.Println("Current hour is", h)

	switch h := time.Now().Hour(); {
	case h >= 18: // 18 to 23
		fmt.Println("good evening")
	case h >= 12: // 12 to 18
		fmt.Println("good afternoon")
	case h >= 6: // 6 to 12
		fmt.Println("good morning")
	default: // 0 to 5
		fmt.Println("good night")
	}

	// h is not available here
	// it's bound to switch statement's block
	// fmt.Println(h)
}
