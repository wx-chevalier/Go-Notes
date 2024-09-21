
package main

import (
	"bufio"
	"os"

	"github.com/inancgumus/learngo/logparser/testing/report"
)

func main() {
	p := report.New()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		p.Parse(in.Text())
	}

	summarize(p.Summarize(), p.Err(), in.Err())
}
