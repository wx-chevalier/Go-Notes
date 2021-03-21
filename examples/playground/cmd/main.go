package main

import (
	"log"
	"ngte/cmd"
)

var name string

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error", err)
	}
}
