package models

type Chat struct {
	ID          uint    `jsonapi:"primary,chats" sql:"unique_index"`
	Name        string  `jsonapi:"attr,name"`
	LastMessage Message `jsonapi:"relation,last_message"`
}
