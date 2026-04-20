
package main

import "fmt"

func main() {
	score := 3

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("low")
	}
}
