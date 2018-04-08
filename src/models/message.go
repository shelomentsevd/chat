package models

import (
	"time"
)

type Message struct {
	ID        uint      `jsonapi:"primary,messages" sql:"unique_index"`
	Content   string    `jsonapi:"attr,content"`
	ChatID    uint      `jsonapi:"attr,chat_id"`
	UserID    uint      `jsonapi:"attr,user_id"`
	CreatedAt time.Time `jsonapi:"attr,created_at"`
	User      User      `jsonapi:"relation,user"`
}
