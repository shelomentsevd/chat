package chats

import (
	"db"
)

func Create(chat *db.Chat) error {
	if err := db.Pool.Create(chat).Error; err != nil {
		return err
	}

	return nil
}

func Get(chat *db.Chat, members bool) error {
	result := db.Pool.Find(chat)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	if members {
		result = result.Related(chat.Members)
	}

	return result.Error
}

func GetList(chats []*db.Chat, limit, offset int) error {
	result := db.Pool.Offset(offset).Limit(limit).Find(&chats)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}
