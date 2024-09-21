
package main

import "fmt"

func main() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else {
		fmt.Println("low")
	}
}
