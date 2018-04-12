package views

import "db"

type Chat struct {
	ID       uint       `jsonapi:"primary,chats"`
	Name     string     `jsonapi:"attr,name"                validate:"required,gte=10"`
	Users    []*User    `jsonapi:"relation,users,omitempty" validate:"required,gte=2,dive"`
	Messages []*Message `jsonapi:"relation,messages,omitempty"`
}

func NewChatView(chat *db.Chat) *Chat {
	var (
		users    []*User
		messages []*Message
	)

	if len(chat.Users) > 0 {
		users = make([]*User, len(chat.Users))
		for i, u := range chat.Users {
			users[i] = NewUserView(u)
		}
	}

	if len(chat.Messages) > 0 {
		messages = make([]*Message, len(chat.Messages))
		for i, m := range chat.Messages {
			messages[i] = NewMessageView(m)
		}
	}

	return &Chat{
		ID:       chat.ID,
		Name:     chat.Name,
		Users:    users,
		Messages: messages,
	}
}
