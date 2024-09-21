
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}
