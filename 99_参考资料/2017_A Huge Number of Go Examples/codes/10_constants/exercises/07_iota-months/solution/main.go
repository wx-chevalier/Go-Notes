
package main

import "fmt"

func main() {
	const (
		Nov = 11 - iota // 11 - 0 = 11
		Oct             // 11 - 1 = 10
		Sep             // 11 - 2 = 9
	)

	fmt.Println(Sep, Oct, Nov)
}
