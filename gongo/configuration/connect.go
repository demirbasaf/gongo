package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, c GongoConfig) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI))

	if err != nil {
		return nil, err
	}

	clientErr := client.Connect(ctx)

	if clientErr != nil {
		return nil, clientErr
	}

	return client, nil
}