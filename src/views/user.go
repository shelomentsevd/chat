package views

import "db"

type User struct {
	ID   uint   `jsonapi:"primary,users" validate:"required"`
	Name string `jsonapi:"attr,name"`
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
