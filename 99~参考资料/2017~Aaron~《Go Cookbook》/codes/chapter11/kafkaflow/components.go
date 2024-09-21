package kafkaflow

import (
	"fmt"
	"strings"

	flow "github.com/trustmaster/goflow"
)

// Upper upper cases the incoming
// stream
type Upper struct {
	flow.Component
	Val <-chan string
	Res chan<- string
}

// OnVal does the encoding then pushes the result onto Res
func (e *Upper) OnVal(val string) {
	e.Res <- strings.ToUpper(val)
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
