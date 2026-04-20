
package main

import "fmt"

func main() {
	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Println(red, blue)
}
