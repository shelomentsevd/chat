package views

import (
	"db"

	"github.com/google/jsonapi"
)

type Message struct {
	ID      uint   `jsonapi:"primary,messages"`
	Content string `jsonapi:"attr,content"`
	User    *User  `jsonapi:"relation,user"`
	meta    *jsonapi.Meta
}

func NewMessageView(message *db.Message, user *User) *Message {
	return &Message{
		ID:      message.ID,
		Content: message.Content,
		User:    user,
		meta: &jsonapi.Meta{
			"created_at": message.CreatedAt,
		},
	}
}

func (message Message) JSONAPIMeta() *jsonapi.Meta {
	return message.meta
}
