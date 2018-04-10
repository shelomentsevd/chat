package chats

import (
	"db"
	"models"
)

func Create(chat *models.Chat) error {
	members := make([]*models.Member, len(chat.Users))

	for i, u := range chat.Users {
		members[i] = &models.Member{
			UserID: u.ID,
		}
	}

	chat.Members = members

	if err := db.Pool.Create(chat).Error; err != nil {
		return err
	}

	return nil
}

func GetByID(chat *models.Chat) error {
	result := db.Pool.Find(chat)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}

func Get(chats []*models.Chat, limit, offset int) error {
	result := db.Pool.Offset(offset).Limit(limit).Find(&chats)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}
