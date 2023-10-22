package database

import (
	"github.com/haakaashs/antino-labs/config"
	"gorm.io/gorm"
)

type UserDb interface {
}

type userDb struct {
	db *gorm.DB
}

func NewUserDb() *userDb {
	return &userDb{
		db: config.Conn,
	}
}
