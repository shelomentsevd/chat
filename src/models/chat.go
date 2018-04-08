package models

type Chat struct {
	ID       uint       `jsonapi:"primary,chats" sql:"unique_index"`
	Name     string     `jsonapi:"attr,name" validate:"required"`
	Users    []*User    `jsonapi:"relation,users,omitempty" validate:"required,gte=2,dive"`
	Messages []*Message `jsonapi:"relation,messages,omitempty"`
	Members  []*Member
}
