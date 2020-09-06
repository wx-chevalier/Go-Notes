
package main

import "fmt"

func main() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}
