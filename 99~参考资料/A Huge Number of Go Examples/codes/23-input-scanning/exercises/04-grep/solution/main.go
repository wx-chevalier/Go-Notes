
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var pattern string
	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}
