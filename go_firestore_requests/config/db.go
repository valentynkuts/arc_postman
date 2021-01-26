package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

const (
	projectID = "copy-of-postmana"
)

var Client *firestore.Client
var Ctx context.Context

func init() {
	Ctx := context.Background()
	var err error
	Client, err = firestore.NewClient(Ctx, projectID)
	if err != nil {
		log.Fatalf("Firestore: %v", err)
	}

}
