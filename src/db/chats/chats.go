package chats

import (
	"db"
)

func Create(chat *db.Chat) error {
	members := make([]*db.Member, len(chat.Users))

	for i, u := range chat.Users {
		members[i] = &db.Member{
			UserID: u.ID,
		}
	}

	chat.Members = members

	if err := db.Pool.Create(chat).Error; err != nil {
		return err
	}

	return nil
}

func GetByID(chat *db.Chat) error {
	result := db.Pool.Find(chat)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}

func Get(chats []*db.Chat, limit, offset int) error {
	result := db.Pool.Offset(offset).Limit(limit).Find(&chats)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}
