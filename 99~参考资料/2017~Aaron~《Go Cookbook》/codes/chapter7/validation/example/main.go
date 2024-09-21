package main

import (
	"fmt"
	"net/http"

	"github.com/agtorre/go-cookbook/chapter7/validation"
)

func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
