package views

import "db"

type User struct {
	ID   uint   `jsonapi:"primary,users"`
	Name string `jsonapi:"attr,name" form:"name" validate:"required" sql:"unique_index"`
}

func NewUserView(user *db.User) *User {
	if user == nil {
		return nil
	}

	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}
