package gongo

import (
	"context"
	"fmt"
	"github.com/afdemirbas/gongo/gongo/configuration"
	"github.com/afdemirbas/gongo/gongo/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

func Connect(c configuration.GongoConfig) error {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.URI))

	if err != nil {
		return err
	}

	return nil
}

func Model(e interface{}) entity.MongoModel {
	et := reflect.TypeOf(e)

	if et.Kind() == reflect.Struct {
		fmt.Println(et.Name())
		val := reflect.Indirect(reflect.ValueOf(e))

		for i := 0; i < val.NumField(); i++ {
			fmt.Println(val.Type().Field(i).Name, val.Field(i).Kind())
		}
	}

	return nil
}