
package main

import "fmt"

func main() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("fast's type is %T\n", fast)

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)
}
