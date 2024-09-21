
package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	// EXERCISE:
	//   Try this alternative (formatted time).

	// fmt.Println(speed, now.Format(time.Kitchen))
}
