
package main

import (
	"fmt"
	"os"
)

// STEPS:
//
// Compile it by typing:
//   go build -o myprogram
//
// Then run it by typing:
//   ./myprogram
//
// If you're on Windows, then type:
//   myprogram

func main() {
	fmt.Println(os.Args[0])
}
