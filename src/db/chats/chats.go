package chats

import (
	"db"
	"models"
)

func Create(chat *models.Chat) error {
	if err := db.Pool.Create(chat).Error; err != nil {
		return err
	}

	return nil
}
