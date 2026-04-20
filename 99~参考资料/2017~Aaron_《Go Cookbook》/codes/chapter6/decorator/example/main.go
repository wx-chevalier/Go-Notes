package main

import "github.com/agtorre/go-cookbook/chapter6/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
