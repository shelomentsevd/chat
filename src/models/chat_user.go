package models

type ChatUser struct {
	ID     uint  `jsonapi:"primary,chat_users" sql:"unique_index"`
	ChatID uint  `jsonapi:"relation,chats"`
	User   *User `jsonapi:"relation,users"`
}
