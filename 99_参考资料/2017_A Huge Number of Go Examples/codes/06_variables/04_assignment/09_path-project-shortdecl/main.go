
package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("css/main.css")

	// or this:
	// dir, file := path.Split("css/main.css")

	fmt.Println("file:", file)
}
