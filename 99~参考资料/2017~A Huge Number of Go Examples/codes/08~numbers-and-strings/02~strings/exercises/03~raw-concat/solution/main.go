
package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
}
