
package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `[command] [string]

Available commands: lower, upper and title`

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	cmd, str := args[1], args[2]
	switch cmd {
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "title":
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	}
}
