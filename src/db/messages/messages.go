package messages

import "db"

func Create(message *db.Message) error {
	if err := db.Pool.Create(message).Error; err != nil {
		return err
	}

	return nil
}

func Get(message *db.Message) error {
	result := db.Pool.Find(message)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}

func GetListByChatID(messages []*db.Message, chat uint, limit, offset int) error {
	result := db.Pool.
		Offset(offset).
		Limit(limit).
		Where(&db.Message{ChatID: chat}).
		Order("created_at desc").
		Find(&messages)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}
