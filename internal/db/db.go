package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	Tasks *mongo.Collection
}

var Coll = Collections{}

var Client *mongo.Database

func initCollections() {
	Coll.Tasks = Client.Collection("tasks")
}

func Connect(connectionString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}

	Client = client.Database("development")
	initCollections()
	return nil
}
