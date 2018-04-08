package models

type Chat struct {
	ID          uint        `jsonapi:"primary,chats" sql:"unique_index"`
	Name        string      `jsonapi:"attr,name"           validate:"required"`
	Users       []*ChatUser `jsonapi:"relation,chat_users" validate:"required,gte=2"`
	LastMessage *Message    `jsonapi:"relation,last_message"`
}
