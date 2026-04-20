package main

import "github.com/agtorre/go-cookbook/chapter6/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
