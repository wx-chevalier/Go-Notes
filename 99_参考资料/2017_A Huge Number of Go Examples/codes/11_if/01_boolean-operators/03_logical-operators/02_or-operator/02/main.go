
package main

import "fmt"

func main() {
	color := "red"

	fmt.Println("reddish colors?",
		// true || false => true (short-circuits)
		color == "red" || color == "dark red",
	)

	color = "dark red"

	fmt.Println("reddish colors?",
		// false || true => true
		color == "red" || color == "dark red",
	)

	fmt.Println("greenish colors?",
		// false || false => false
		color == "green" || color == "light green",
	)
}
