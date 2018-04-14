package db

import "github.com/jinzhu/gorm"

type Option struct {
	apply func(*gorm.DB) *gorm.DB
}

func WithOffset(offset int) Option {
	return Option{
		apply: func(db *gorm.DB) *gorm.DB {
			return db.Offset(offset)
		},
	}
}

func WithLimit(limit int) Option {
	return Option{
		apply: func(db *gorm.DB) *gorm.DB {
			return db.Limit(limit)
		},
	}
}

func WithIDs(ids ...uint) Option {
	return Option{
		apply: func(db *gorm.DB) *gorm.DB {
			return db.Where(ids)
		},
	}
}

func WithOrder(order string) Option {
	return Option{
		apply: func(db *gorm.DB) *gorm.DB {
			return db.Order(order)
		},
	}
}

func WithCondition(model interface{}) Option {
	return Option{
		apply: func(db *gorm.DB) *gorm.DB {
			return db.Where(model)
		},
	}
}
