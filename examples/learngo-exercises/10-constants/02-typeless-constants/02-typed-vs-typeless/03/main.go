
package main

import "fmt"

func main() {
	const min = 42

	var f float64

	f = min // OK when min is typeless

	fmt.Println(f)
}
