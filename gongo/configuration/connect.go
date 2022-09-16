package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, c GongoConfig) (*GongoConnection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.URI))

	gongoConnection := GongoConnection{
		Client:   client,
		Database: c.Database,
	}

	if err != nil {
		return nil, err
	}

	clientErr := client.Connect(ctx)

	if clientErr != nil {
		return nil, clientErr
	}

	return &gongoConnection, nil
}