package models

type User struct {
	ID       uint   `jsonapi:"primary,users" sql:"unique_index"`
	Name     string `jsonapi:"attr,name" form:"name" validate:"required" sql:"unique_index"`
	Password string `form:"password" validate:"required"`
}
