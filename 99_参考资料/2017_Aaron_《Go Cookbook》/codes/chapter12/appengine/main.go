package main

import (
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func main() {
	ctx := context.Background()
	log.SetOutput(os.Stderr)

	// Set this in app.yaml when running in production.
	projectID := os.Getenv("GCLOUD_DATASET_ID")

	datastoreClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	c := Controller{datastoreClient}

	http.HandleFunc("/", c.handle)
	appengine.Main()
}
