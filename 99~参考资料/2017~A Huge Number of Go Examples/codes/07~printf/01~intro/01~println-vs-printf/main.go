
package main

import "fmt"

func main() {
	ops, ok, fail := 2350, 543, 433

	fmt.Println(
		"total:", ops, "- success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)
}
