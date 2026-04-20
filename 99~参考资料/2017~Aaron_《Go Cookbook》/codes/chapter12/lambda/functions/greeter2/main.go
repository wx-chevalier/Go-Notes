package main

import (
	"encoding/json"
	"fmt"

	"github.com/apex/go-apex"
)

type message struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		resp := map[string]string{
			"greeting": fmt.Sprintf("Hello, %s %s", m.FirstName, m.LastName),
		}

		return resp, nil
	})
}
