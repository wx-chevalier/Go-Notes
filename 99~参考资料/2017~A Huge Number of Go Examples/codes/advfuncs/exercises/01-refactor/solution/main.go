
package main

import "fmt"

func main() {
	fmt.Println(headerOf("gif"))
}

var headers = map[string]string{
	"png": "\x89PNG\r\n\x1a\n",
	"jpg": "\xff\xd8\xff",
}

func headerOf(format string) (header string) {
	defer func() {
		if header == "" {
			panic("unknown format: " + format)
		}
	}()
	return headers[format]
}
