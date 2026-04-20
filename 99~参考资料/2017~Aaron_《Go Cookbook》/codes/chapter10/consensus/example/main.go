package main

import (
	"net/http"

	"github.com/agtorre/go-cookbook/chapter10/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
