package db

import "time"

type Chat struct {
	ID       uint `sql:"unique_index"`
	Name     string
	Messages []*Message
	Members  []*Member
}

type Member struct {
	ID     uint `sql:"unique_index"`
	UserID uint
	ChatID uint
}

type Message struct {
	ID        uint `sql:"unique_index"`
	Content   string
	ChatID    uint
	UserID    uint
	CreatedAt time.Time
	User      *User
}

type User struct {
	ID       uint   `sql:"unique_index"`
	Name     string `sql:"unique_index"`
	Password string
}
