package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
