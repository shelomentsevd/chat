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

func Get(chats []*models.Chat, limit, offset int) error {
	if err := db.Pool.
		Offset(offset).
		Limit(limit).
		Find(&chats).Error; err != nil {
		return err
	}

	return nil
}
