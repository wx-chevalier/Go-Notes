
package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Remove the Magic
//
//  Get rid of the magic numbers in the following code.
//
// RESTRICTIONS
//  1. You should declare 3 constants named:
//       hoursInDay, daysInWeek, and hoursPerWeek
//
//  2. And, hoursPerWeek constant should be initialized
//     using hoursInDay and daysInWeek constants.
//
// EXPECTED OUTPUT
//  There are 840 hours in 5 weeks.
// ---------------------------------------------------------

func main() {
	fmt.Printf("There are %d hours in %d weeks.\n",
		24*7*5, 5)
}
