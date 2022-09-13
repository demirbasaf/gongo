package examples

type User struct {
	Name string `json:"name"`
}

type Article struct {
	Title      string `json:"title,omitempty"`
	Author     User   `json:"author"`
	Users      []User `json:"users,omitempty"`
	IsArchived bool   `json:"is_archived,omitempty"`
}