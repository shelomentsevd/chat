package users

import (
	"db"
	"models"
)

func GetByName(name string) (*models.User, error) {
	var user models.User

	if err := db.Pool.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
