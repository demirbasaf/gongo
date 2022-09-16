package entity

import (
	"context"
	"fmt"
)

func New(ctx context.Context, m *Model, e interface{}) interface{} {
	collection := m.Client.Database(m.Database).Collection(m.Collection)

	fmt.Println(collection.Name())

	result, err := collection.InsertOne(ctx, e)

	if err != nil {
		fmt.Println("Error")
	}

	return result.InsertedID
}