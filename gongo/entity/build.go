package entity

import (
	"github.com/afdemirbas/gongo/gongo/configuration"
	"reflect"
)

func NewModel(c string, e interface{}, gc *configuration.GongoConnection) MongoModel {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Struct {
		name := t.Name()
		fields := extractFields(e)

		return &Model{
			Name:       name,
			Fields:     fields,
			Client:     *gc.Client,
			Database:   gc.Database,
			Collection: c,
		}
	}
	return nil
}

func extractFields(e interface{}) []string {
	var fields []string
	val := reflect.Indirect(reflect.ValueOf(e))

	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	return fields
}