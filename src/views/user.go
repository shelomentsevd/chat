package views

import "db"

type User struct {
	ID   uint   `jsonapi:"primary,users" sql:"unique_index"`
	Name string `jsonapi:"attr,name" form:"name" validate:"required" sql:"unique_index"`
	//Password string `form:"password" validate:"required"`
}

func NewUserView(user *db.User) *User {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}
