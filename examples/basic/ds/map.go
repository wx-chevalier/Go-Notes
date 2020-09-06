package main

import (
	"fmt"
)

var m = make(map[string]string)

func main() {
	if v, ok := m["key"]; ok {
		fmt.Print(v)
	}

}
