package main

import (
	"fmt"

	"github.com/agtorre/go-cookbook/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
