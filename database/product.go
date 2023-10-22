package database

import (
	"github.com/haakaashs/antino-labs/config"
	"gorm.io/gorm"
)

type ProductDb interface {
}

type productDb struct {
	db *gorm.DB
}

func NewProductDb() *productDb {
	return &productDb{
		db: config.Conn,
	}
}
