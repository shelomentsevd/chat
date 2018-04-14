package db

import (
	"time"
)

type Chat struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"not null"`
	Members []*Member
}

type Member struct {
	ID     uint `gorm:"primary_key"`
	UserID uint
	ChatID uint
}

type Message struct {
	ID        uint   `gorm:"primary_key"`
	Content   string `gorm:"not null"`
	ChatID    uint
	UserID    uint
	CreatedAt time.Time
	User      *User
}

type User struct {
	ID       uint    `gorm:"primary_key"`
	Name     string  `sql:"unique_index" gorm:"not null"`
	Password *string `gorm:"not null"`
}
