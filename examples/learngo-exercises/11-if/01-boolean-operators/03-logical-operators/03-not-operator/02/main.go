
package main

import (
	"fmt"
)

func main() {
	var on bool

	on = !on
	fmt.Println(on)

	on = !!on
	fmt.Println(on)
}
