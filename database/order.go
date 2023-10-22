package database

import (
	"github.com/haakaashs/antino-labs/config"
	"gorm.io/gorm"
)

type OrderDb interface {
}

type orderDb struct {
	db *gorm.DB
}

func NewOrderDb() *orderDb {
	return &orderDb{
		db: config.Conn,
	}
}
