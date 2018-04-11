package views

import (
	"db"

	"github.com/google/jsonapi"
)

type Message struct {
	ID      uint   `jsonapi:"primary,messages" sql:"unique_index"`
	Content string `jsonapi:"attr,content"`
	User    *User  `jsonapi:"relation,user"`
	meta    *jsonapi.Meta
}

func NewMessageView(message *db.Message) *Message {
	return &Message{
		ID:      message.ID,
		Content: message.Content,
		User:    NewUserView(message.User),
		meta: &jsonapi.Meta{
			"created_at": message.CreatedAt,
		},
	}
}

func (message Message) JSONAPIMeta() *jsonapi.Meta {
	return message.meta
}
