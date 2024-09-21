package main

import "github.com/agtorre/go-cookbook/chapter5/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
