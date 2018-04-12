package views

import "db"

type Chat struct {
	ID       uint       `jsonapi:"primary,chats"`
	Name     string     `jsonapi:"attr,name"                validate:"required,gte=10"`
	Users    []*User    `jsonapi:"relation,users,omitempty" validate:"required,gte=1,dive"`
	Messages []*Message `jsonapi:"relation,messages,omitempty"`
}

func NewChatView(chat *db.Chat, users []*User, messages []*Message) *Chat {
	return &Chat{
		ID:       chat.ID,
		Name:     chat.Name,
		Users:    users,
		Messages: messages,
	}
}
