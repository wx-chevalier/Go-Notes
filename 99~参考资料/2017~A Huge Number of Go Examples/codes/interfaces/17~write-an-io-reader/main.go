
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// resp, err := http.Get("https://inancgumus.github.com/x/rosie.jpg")
	resp, err := http.Get("https://inancgumus.github.com/x/rosie.unknown")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create("rosie.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	n, err := io.Copy(file, pngReader(resp.Body))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%d bytes transferred.\n", n)
}
