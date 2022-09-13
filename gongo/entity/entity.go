package entity

type MongoModel interface {
	New()
	Find()
	FindMany()
	FindOne()
	Update()
	Delete()
}

type Entity struct{}