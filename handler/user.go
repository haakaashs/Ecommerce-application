package handler

import (
	"log"
	"net/http"

	"github.com/haakaashs/antino-labs/service"
)

type UserHandler interface {
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	funcdesc := "CreateUser"
	log.Println("enter " + funcdesc)

	log.Println("exit " + funcdesc)
}
