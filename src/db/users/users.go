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

func GetList(users []*db.User, limit, offset int) error {
	result := db.Pool.Limit(limit).Offset(offset).Find(&users)

	return result.Error
}

func GetByIDs(ids ...uint) ([]*db.User, error) {
	users := make([]*db.User, 0)

	result := db.Pool.Where(ids).Find(&users)

	if result.RecordNotFound() {
		return nil, db.RecordNotFound
	}

	return users, nil
}

func Create(user *db.User) error {
	return db.Pool.Create(user).Error
}
