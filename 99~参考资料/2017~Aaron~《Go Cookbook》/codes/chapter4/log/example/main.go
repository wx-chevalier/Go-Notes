package main

import (
	"fmt"

	"github.com/agtorre/go-cookbook/chapter4/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
