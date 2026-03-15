
package record

import "encoding/json"

// UnmarshalJSON to a record.
func (r *Record) UnmarshalJSON(data []byte) error {
	type rjson Record

	err := json.Unmarshal(data, (*rjson)(r))
	if err != nil {
		return err
	}

	return r.validate()
}
