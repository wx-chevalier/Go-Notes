package main

import (
	"fmt"
	"log"

	"github.com/agtorre/go-cookbook/chapter12/firebase"
)

func main() {
	f, err := firebase.Authenticate()
	if err != nil {
		log.Fatalf("error authenticating")
	}
	f.Set("key", []string{"val1", "val2"})
	res, _ := f.Get()
	fmt.Println(res)

	vals := res["key"].([]interface{})
	vals = append(vals, map[string][]string{"key2": []string{"val3"}})
	f.Set("key", vals)
	res, _ = f.Get()
	fmt.Println(res)
}
