
package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	if bytes.Equal(png, header) {
		fmt.Println("They are equal")
	}
}
