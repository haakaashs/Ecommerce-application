package handler

import "github.com/haakaashs/antino-labs/service"

type UserHandle interface {
}

type userHandle struct {
	userService service.UserService
}

func NewUserHandle(userService service.UserService) *userHandle {
	return &userHandle{
		userService: userService,
	}
}
