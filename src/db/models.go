package db

import "time"

type Chat struct {
	ID       uint       `jsonapi:"primary,chats" sql:"unique_index"`
	Name     string     `jsonapi:"attr,name" validate:"required"`
	Users    []*User    `jsonapi:"relation,users,omitempty" validate:"required,gte=2,dive"`
	Messages []*Message `jsonapi:"relation,messages,omitempty"`
	Members  []*Member
}

type Member struct {
	ID     uint `sql:"unique_index"`
	UserID uint
	ChatID uint
}

type Message struct {
	ID        uint      `jsonapi:"primary,messages" sql:"unique_index"`
	Content   string    `jsonapi:"attr,content"`
	ChatID    uint      `jsonapi:"attr,chat_id"`
	UserID    uint      `jsonapi:"attr,user_id"`
	CreatedAt time.Time `jsonapi:"attr,created_at"`
	User      *User     `jsonapi:"relation,user"`
}

type User struct {
	ID       uint   `jsonapi:"primary,users" sql:"unique_index"`
	Name     string `jsonapi:"attr,name" form:"name" validate:"required" sql:"unique_index"`
	Password string `form:"password" validate:"required"`
}
