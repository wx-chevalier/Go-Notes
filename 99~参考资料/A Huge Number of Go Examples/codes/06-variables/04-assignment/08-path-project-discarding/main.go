
package main

import (
	"fmt"
	"path"
)

func main() {
	var file string

	_, file = path.Split("css/main.css")

	fmt.Println("file:", file)
}
