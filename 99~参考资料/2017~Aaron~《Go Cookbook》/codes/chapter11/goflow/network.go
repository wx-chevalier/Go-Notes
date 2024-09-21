package goflow

import flow "github.com/trustmaster/goflow"

// EncodingApp creates a flow-based
// pipeline to encode and print the
// result
type EncodingApp struct {
	flow.Graph
}

// NewEncodingApp wires together the componets
func NewEncodingApp() *EncodingApp {
	e := &EncodingApp{}
	e.InitGraphState()

	// define component types
	e.Add(&Encoder{}, "encoder")
	e.Add(&Printer{}, "printer")

	// connect the components using channels
	e.Connect("encoder", "Res", "printer", "Line")

	// map the in channel to Val, which is
	// tied to OnVal function
	e.MapInPort("In", "encoder", "Val")

	return e
}
