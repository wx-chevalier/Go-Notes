package main

import "github.com/agtorre/go-cookbook/chapter5/mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
