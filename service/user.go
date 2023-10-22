package service

import (
	"errors"
	"log"

	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/utils"
)

type UserService interface {
	CreateUser(models.User) (uint, error)
	GetUserById(uint64) (models.User, error)
	GetUsers() ([]models.User, error)
	UserLogin(resources.UserCredential) error
	DeleteUserById(uint64) error
}

type userService struct {
	userDB database.UserDb
}

func NewUserService(userDB database.UserDb) *userService {
	return &userService{
		userDB: userDB,
	}
}

func (s *userService) CreateUser(user models.User) (uint, error) {
	funcdesc := "CreateUser"
	log.Println("enter service" + funcdesc)

	if !utils.IsEmailValid(user.Email) {
		return 0, errors.New("email validation failed")
	}

	userId, err := s.userDB.CreateUser(user)
	if err != nil {
		return userId, err
	}

	log.Println("exit " + funcdesc)
	return userId, nil
}

func (s *userService) GetUserById(userId uint64) (user models.User, err error) {
	funcdesc := "GetUserById"
	log.Println("enter service" + funcdesc)

	user, err = s.userDB.GetUserById(userId)
	if err != nil {
		return user, err
	}

	log.Println("exit " + funcdesc)
	return user, nil
}

func (s *userService) GetUsers() (users []models.User, err error) {
	funcdesc := "GetUsers"
	log.Println("enter service" + funcdesc)

	users, err = s.userDB.GetUsers()
	if err != nil {
		return users, err
	}

	log.Println("exit " + funcdesc)
	return users, nil
}

func (s *userService) UserLogin(input resources.UserCredential) (err error) {
	funcdesc := "UserLogin"
	log.Println("enter service" + funcdesc)

	user, err := s.userDB.UserLogin(input)
	if err != nil {
		return err
	} else if user.Password != input.Password {
		return errors.New("incorrect Password")
	}

	log.Println("exit " + funcdesc)
	return nil
}

func (s *userService) DeleteUserById(userId uint64) (err error) {
	funcdesc := "DeleteUserById"
	log.Println("enter service" + funcdesc)

	err = s.userDB.DeleteUserById(userId)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
