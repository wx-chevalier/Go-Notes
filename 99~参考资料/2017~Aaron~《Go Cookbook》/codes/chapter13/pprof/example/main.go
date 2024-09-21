package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/agtorre/go-cookbook/chapter13/pprof/crypto"
)

func main() {

	http.HandleFunc("/guess", crypto.GuessHandler)
	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
