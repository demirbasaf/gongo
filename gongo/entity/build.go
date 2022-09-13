package entity

import "reflect"

func NewModel(e interface{}) MongoModel {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Struct {
		name := t.Name()
		fields := extractFields(e)

		return &Model{
			Name:   name,
			Fields: fields,
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