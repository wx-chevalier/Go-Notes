package main

import "github.com/agtorre/go-cookbook/chapter1/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
