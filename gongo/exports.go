package gongo

import (
	"context"
	"github.com/afdemirbas/gongo/gongo/configuration"
	"github.com/afdemirbas/gongo/gongo/entity"
)

func Connect(ctx context.Context, c configuration.GongoConfig) (*configuration.GongoConnection, error) {
	return configuration.Connect(ctx, c)
}

func Model(c string, e interface{}, gc *configuration.GongoConnection) entity.MongoModel {
	return entity.NewModel(c, e, gc)
}