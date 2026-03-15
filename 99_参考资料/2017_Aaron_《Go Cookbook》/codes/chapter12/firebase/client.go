package firebase

import (
	"log"

	"gopkg.in/zabawaba99/firego.v1"
)

// Client Interface for mocking
type Client interface {
	Get() (map[string]interface{}, error)
	Set(key string, value interface{}) error
}
type firebaseClient struct {
	*firego.Firebase
}

func (f *firebaseClient) Get() (map[string]interface{}, error) {
	var v2 map[string]interface{}
	if err := f.Value(&v2); err != nil {
		log.Fatalf("error getting")
	}
	return v2, nil
}

func (f *firebaseClient) Set(key string, value interface{}) error {
	v := map[string]interface{}{key: value}
	if err := f.Firebase.Set(v); err != nil {
		return err
	}
	return nil
}
