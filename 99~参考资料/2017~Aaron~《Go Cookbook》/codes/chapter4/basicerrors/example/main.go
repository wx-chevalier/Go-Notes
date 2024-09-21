package main

import (
	"fmt"

	"github.com/agtorre/go-cookbook/chapter4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
