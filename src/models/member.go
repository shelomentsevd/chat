package models

type Member struct {
	ID     uint `sql:"unique_index"`
	UserID uint
	ChatID uint
}
