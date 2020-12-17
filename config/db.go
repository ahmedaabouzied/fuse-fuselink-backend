package config

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToMongoDB(ctx context.Context, connectionString string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to MongoDB")
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to MongoDB")
	}
	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to MongoDB")
	}
	log.Println("Connected to MongoDB cluster")
	return client, nil
}
