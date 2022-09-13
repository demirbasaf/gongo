package gongo

import (
	"context"
	"github.com/afdemirbas/gongo/gongo/configuration"
	"github.com/afdemirbas/gongo/gongo/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(c configuration.GongoConfig) error {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.URI))

	if err != nil {
		return err
	}

	return nil
}

func Model(e interface{}) entity.MongoModel {
	return entity.NewModel(e)
}