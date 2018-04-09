package chats

import (
	"db"
	"models"
)

// TODO: RecordNotFound

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
	if err := db.Pool.Find(chat).Error; err != nil {
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
