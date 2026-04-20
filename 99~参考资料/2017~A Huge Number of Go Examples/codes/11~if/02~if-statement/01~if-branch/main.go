
package main

import "fmt"

func main() {
	score, valid := 5, true

	if score > 3 && valid {
		fmt.Println("good")
	}
}
