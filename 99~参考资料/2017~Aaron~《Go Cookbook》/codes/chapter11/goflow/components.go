package goflow

import (
	"encoding/base64"
	"fmt"

	flow "github.com/trustmaster/goflow"
)

// Encoder base64 encodes all input
type Encoder struct {
	flow.Component
	Val <-chan string
	Res chan<- string
}

// OnVal does the encoding then pushes the result onto Res
func (e *Encoder) OnVal(val string) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	e.Res <- fmt.Sprintf("%s => %s", val, encoded)
}

// Printer is a component for printing to stdout
type Printer struct {
	flow.Component
	Line <-chan string
}

// OnLine Prints the current line received
func (p *Printer) OnLine(line string) {
	fmt.Println(line)
}
