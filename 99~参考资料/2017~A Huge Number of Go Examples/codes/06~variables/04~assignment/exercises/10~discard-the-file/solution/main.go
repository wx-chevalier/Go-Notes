
package main

import (
	"fmt"
	"path"
)

func main() {
	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)
}
