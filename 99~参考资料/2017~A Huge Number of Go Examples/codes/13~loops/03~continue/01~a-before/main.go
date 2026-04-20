
package main

import (
	"fmt"
)

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			// this continue creates an endless loop
			// since the code never increases the i
			continue
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}
