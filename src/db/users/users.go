package users

import (
	"db"
)

func Get(user *db.User) error {
	result := db.Pool.Find(user)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}

func Create(user *db.User) error {
	return db.Pool.Create(user).Error
}
