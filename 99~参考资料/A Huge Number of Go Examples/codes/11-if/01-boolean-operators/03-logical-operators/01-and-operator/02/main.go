
package main

import "fmt"

func main() {
	speed := 100
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
	)

	speed = 20
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
		// ^- short-circuits in the first expression here
		//    because, it becomes false
	)

	speed = 50
	fmt.Println("within limits?",
		speed >= 40 && speed <= 55,
	)

	// ERROR: invalid
	// both operands should be booleans
	// 1 && 2
	fmt.Println(1 == 1 && 2 == 2)
}
