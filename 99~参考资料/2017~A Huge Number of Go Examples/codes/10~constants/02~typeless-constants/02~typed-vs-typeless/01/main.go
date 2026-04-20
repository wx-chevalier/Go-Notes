
package main

import "fmt"

func main() {
	const min int = 42

	var i int
	i = min // OK

	fmt.Println(i)
}
