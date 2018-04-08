package models

type Chat struct {
	ID       uint       `jsonapi:"primary,chats" sql:"unique_index"`
	Name     string     `jsonapi:"attr,name"        validate:"required"`
	Members  []*Member  `jsonapi:"relation,members" validate:"required,gte=2,dive"`
	Messages []*Message `jsonapi:"relation,messages,omitempty"`
}
