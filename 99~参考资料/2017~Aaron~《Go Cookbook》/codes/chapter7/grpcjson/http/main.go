package main

import (
	"fmt"
	"net/http"

	"github.com/agtorre/go-cookbook/chapter7/grpcjson/internal"
)

func main() {
	c := Controller{KeyValue: internal.NewKeyValue()}
	http.HandleFunc("/set", c.SetHandler)
	http.HandleFunc("/get", c.GetHandler)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
