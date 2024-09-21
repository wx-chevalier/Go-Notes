
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.NumCPU() is a call expression
	fmt.Println(runtime.NumCPU() + 1)
}
