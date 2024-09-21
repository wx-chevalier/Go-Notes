package kafkaflow

import flow "github.com/trustmaster/goflow"

// UpperApp creates a flow-based
// pipeline to upper case and print the
// result
type UpperApp struct {
	flow.Graph
}

// NewUpperApp wires together the compoents
func NewUpperApp() *UpperApp {
	u := &UpperApp{}
	u.InitGraphState()

	u.Add(&Upper{}, "upper")
	u.Add(&Printer{}, "printer")

	u.Connect("upper", "Res", "printer", "Line")
	u.MapInPort("In", "upper", "Val")

	return u
}
