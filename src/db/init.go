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

	return err
}
