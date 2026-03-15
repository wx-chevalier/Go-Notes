
package main

import (
	"log"
	"os"
)

func main() {
	// p := pipe.Default(
	// 	os.Stdin, os.Stdout,
	// 	pipe.FilterBy(pipe.DomainExtFilter("com", "io")),
	// 	pipe.GroupBy(pipe.DomainGrouper),
	// )

	p, err := fromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
