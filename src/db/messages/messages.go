package messages

import "db"

func Create(message *db.Message) error {
	if err := db.Pool.Create(message).Error; err != nil {
		return err
	}

	return nil
}
