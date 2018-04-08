package models

type Member struct {
	ID     uint `jsonapi:"primary,members" sql:"unique_index"`
	UserID uint `jsonapi:"attr,user_id" validate:"required,ne=0"`
	ChatID uint `jsonapi:"attr,chat_id"`
}
