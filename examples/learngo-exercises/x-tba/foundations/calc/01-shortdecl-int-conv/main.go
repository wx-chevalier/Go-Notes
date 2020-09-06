
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1.
	// var a int
	// var b int

	// 2.
	// var (
	//   a int
	//   b int
	// )

	// 3.
	// var a, b int

	// 1.
	// fmt.Println(a, "+", b, "=", a+b)

	// 2.
	// fmt.Printf("%v + %v = %v\n", a, b, a+b)

	// ----
	// lesson: multi-return funcs, %v, and _

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
