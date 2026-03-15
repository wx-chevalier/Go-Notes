package main

import (
	"encoding/json"
	"os"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Input takes in a secret
type Input struct {
	Secret string `json:"secret"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.SetHandler(text.New(os.Stderr))

		var input Input
		if err := json.Unmarshal(event, &input); err != nil {
			log.WithError(err).Error("failed to unmarshal key input")
			return nil, err
		}
		log.WithField("secret", input.Secret).Info("secret guessed")

		if input.Secret == "klaatu barada nikto" {
			return "secret guessed!", nil
		}
		return "try again", nil
	})
}
