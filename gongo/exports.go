package gongo

import (
	"context"
	"github.com/afdemirbas/gongo/gongo/configuration"
	"github.com/afdemirbas/gongo/gongo/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func Connect(ctx context.Context, c configuration.GongoConfig) (*mongo.Client, error) {
	return configuration.Connect(ctx, c)
}

func Model(e interface{}) entity.MongoModel {
	return entity.NewModel(e)
}