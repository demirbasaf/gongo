package configuration

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	URI      string
	Password string
}

func (c *Connection) Connect() {
	mongo.Connect(context.TODO(), options.Client().ApplyURI(c.URI))
}

func BuildConnection(uri string, pw string) *Connection {
	return &Connection{
		URI:      uri,
		Password: pw,
	}
}