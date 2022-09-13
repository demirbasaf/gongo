package entity

type MongoModel interface {
	New(e interface{}) string
	Find()
	FindMany()
	FindOne()
	Update()
	Delete()
}

type Model struct {
	Name   string
	Fields []string
}

func (m *Model) New(e interface{}) string {
	return New(*m)
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