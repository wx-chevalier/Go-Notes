package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go sub", flag.ContinueOnError)
		goCmd.StringVar(&name, "name", "Go", "help")
		_ = goCmd.Parse(args[1:])
	}

	log.Printf(name)

}
