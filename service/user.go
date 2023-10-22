package service

import "github.com/haakaashs/antino-labs/database"

type UserService interface {
}

type userService struct {
	userDB database.UserDb
}

func NewUserService(userDB database.UserDb) *userService {
	return &userService{
		userDB: userDB,
	}
}
