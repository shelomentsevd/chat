package models

import "github.com/m4rw3r/uuid"

type User struct {
	ID       uuid.UUID `json:"id"   sql:"unique_index"`
	Name     string    `json:"name" sql:"unique_index"`
	Password string    `json:"-"`
}
