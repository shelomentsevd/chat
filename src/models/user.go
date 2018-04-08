package models

type User struct {
	ID       uint   `jsonapi:"primary,users" sql:"unique_index"`
	Name     string `jsonapi:"attr,name"     sql:"unique_index"`
	Password string
}
