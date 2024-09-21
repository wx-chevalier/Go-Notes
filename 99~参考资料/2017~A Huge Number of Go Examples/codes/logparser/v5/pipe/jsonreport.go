
package pipe

import (
	"encoding/json"
	"io"
)

// JSONReport generates a JSON report.
type JSONReport struct {
	w io.Writer
}

// NewJSONReport returns a JSON report generator.
func NewJSONReport(w io.Writer) *JSONReport {
	return &JSONReport{w: w}
}

// Consume the records and generate a JSON report.
func (t *JSONReport) Consume(records Iterator) error {
	enc := json.NewEncoder(t.w)

	encode := func(r Record) error {
		return enc.Encode(&r)
	}

	return records.Each(encode)
}
