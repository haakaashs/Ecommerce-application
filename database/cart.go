package database

import (
	"github.com/haakaashs/antino-labs/config"
	"gorm.io/gorm"
)

type CartDB interface {
}

type cartDB struct {
	db *gorm.DB
}

func NewCartDB() *cartDB {
	return &cartDB{
		db: config.Conn,
	}
}
