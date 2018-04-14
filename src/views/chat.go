package views

import "db"

type Chat struct {
	ID    uint    `jsonapi:"primary,chats"`
	Name  string  `jsonapi:"attr,name"                validate:"required,gte=10"`
	Users []*User `jsonapi:"relation,users,omitempty" validate:"required,gte=1,dive"`
}

func NewChatView(chat *db.Chat, users []*User) *Chat {
	if chat == nil {
		return nil
	}

	return &Chat{
		ID:    chat.ID,
		Name:  chat.Name,
		Users: users,
	}
}
