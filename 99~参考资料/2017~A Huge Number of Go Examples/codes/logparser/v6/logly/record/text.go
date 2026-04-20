
package record

import (
	"fmt"
	"strconv"
	"strings"
)

// FromText unmarshals the log line into this record.
func (r *Record) FromText(p []byte) (err error) {
	fields := strings.Fields(string(p))

	if len(fields) != fieldsLength {
		return fmt.Errorf("wrong number of fields %q", fields)
	}

	r.Domain = fields[0]
	r.Page = fields[1]

	const msg = "record.UnmarshalText %q: %v"
	if r.Visits, err = strconv.Atoi(fields[2]); err != nil {
		return fmt.Errorf(msg, "visits", err)
	}
	if r.Uniques, err = strconv.Atoi(fields[3]); err != nil {
		return fmt.Errorf(msg, "uniques", err)
	}

	return r.validate()
}
