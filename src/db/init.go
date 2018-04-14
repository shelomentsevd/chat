package db

import (
	"time"

	"errors"

	"github.com/jinzhu/gorm"
)

var RecordNotFound = errors.New("record not found")

var Pool *gorm.DB

func Init(params string, maxConnections, maxIdleConnections int, lifeTime time.Duration) (err error) {
	if Pool, err = gorm.Open("postgres", params); err != nil {
		return
	}

	Pool.DB().SetMaxIdleConns(maxIdleConnections)
	Pool.DB().SetConnMaxLifetime(lifeTime)
	Pool.DB().SetMaxOpenConns(maxConnections)

	if err = Pool.DB().Ping(); err != nil {
		return
	}

	tables := []interface{}{
		&User{},
		&Chat{},
		&Message{},
		&Member{},
	}

	for _, table := range tables {
		var result *gorm.DB
		if !Pool.HasTable(table) {
			result = Pool.CreateTable(table)
		} else {
			result = Pool.AutoMigrate(table)
		}

		if result.Error != nil {
			return err
		}
	}

	// Foreign keys for messages table
	if err = Pool.Model(&Message{}).AddForeignKey("chat_id", "chats(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return
	}

	if err = Pool.Model(&Message{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return
	}

	// Foreign keys for members table
	if err = Pool.Model(&Member{}).AddForeignKey("chat_id", "chats(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return
	}
	if err = Pool.Model(&Member{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return
	}

	// Unique index for pair user_id and chat_id in members table
	if err = Pool.Exec("CREATE UNIQUE INDEX IF NOT EXISTS user_id_chat_id_idx ON members(user_id,chat_id)").Error; err != nil {
		return
	}

	return
}
