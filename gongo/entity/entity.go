package entity

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoModel interface {
	New(ctx context.Context, e interface{}) interface{}
	Find()
	FindMany()
	FindOne()
	Update()
	Delete()
}

type Model struct {
	Database   string
	Collection string
	Client     mongo.Client
	Name       string
	Fields     []string
}

func (m *Model) New(ctx context.Context, e interface{}) interface{} {
	return New(ctx, m, e)
}

func (m *Model) Find() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) FindMany() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) FindOne() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Update() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Delete() {
	//TODO implement me
	panic("implement me")
}