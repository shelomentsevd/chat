package users

import (
	"db"
	"models"
)

func Get(user *models.User) error {
	result := db.Pool.Find(user)

	if result.RecordNotFound() {
		return db.RecordNotFound
	}

	return result.Error
}

func Create(user *models.User) error {
	return db.Pool.Create(user).Error
}
