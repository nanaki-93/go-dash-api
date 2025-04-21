package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"log"
	"os"
)

type FireDB struct {
	*firestore.Client
}

func NewFirebaseConnection() *firestore.Client {

	home, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "/internal/resources/firebase-config.json")

	client, err := firestore.NewClient(ctx, "go-dash-api", opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	return client
}
