package models

type User struct {
	ID       uint   `                                    jsonapi:"primary,users" sql:"unique_index"`
	Name     string `form:"name"     validate:"required" jsonapi:"attr,name"     sql:"unique_index"`
	Password string `form:"password" validate:"required"`
}
