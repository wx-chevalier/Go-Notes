
package main

import (
	"encoding/json"
	"fmt"
)

// Wizard is one of the greatest of people
type Wizard struct {
	// name won't be marshalled (should be exported)
	Name     string `json:"name,omitempty"`
	Lastname string `json:"last_name"`
	Nick     string `json:"-"`
}

func main() {
	wizards := []Wizard{
		{Name: "Albert", Lastname: "Einstein", Nick: "emc2"},
		{Name: "Isaac", Lastname: "Newton", Nick: "apple"},
		{Name: "Stephen", Lastname: "Hawking", Nick: "blackhole"},
		{Name: "Marie", Lastname: "Curie", Nick: "radium"},
		{Name: "", Lastname: "Darwin", Nick: "fittest"},
	}

	bytes, err := json.MarshalIndent(wizards, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
