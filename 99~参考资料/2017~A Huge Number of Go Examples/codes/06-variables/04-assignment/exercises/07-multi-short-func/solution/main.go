
package main

import (
	"fmt"
)

func main() {
	_, b := multi()

	fmt.Println(b)
}

func multi() (int, int) {
	return 5, 4
}
