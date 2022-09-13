package entity

import (
	"fmt"
	"reflect"
)

type MongoEntity interface {
	Create()
	Find()
	FindOne()
	Update()
	Delete()
}

type Entity struct{}

func NewModel(entity interface{}) MongoEntity {
	entityType := reflect.TypeOf(entity)

	if entityType.Kind() == reflect.Struct {
		fmt.Println(entityType.Name())
		val := reflect.Indirect(reflect.ValueOf(entity))

		for i := 0; i < val.NumField(); i++ {
			fmt.Println(val.Type().Field(i).Name, val.Field(i).Kind())
		}
	}

	return nil
}