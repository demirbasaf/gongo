package configuration

type Connection struct {
	URI      string
	Password string
}

func (c *Connection) Connect() {}

func BuildConnection(uri string, pw string) *Connection {
	return &Connection{
		URI:      uri,
		Password: pw,
	}
}