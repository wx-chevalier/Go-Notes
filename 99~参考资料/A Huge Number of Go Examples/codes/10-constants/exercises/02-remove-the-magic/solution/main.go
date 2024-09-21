
package main

import "fmt"

func main() {
	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)

	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)
}
