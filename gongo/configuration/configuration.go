package configuration

import "go.mongodb.org/mongo-driver/mongo"

type GongoConfig struct {
	Database string
	URI      string
}

type GongoConnection struct {
	Client   *mongo.Client
	Database string
}