
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fmt.Println(strings.ToUpper(in.Text()))
	}
}
