package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"gorm.io/gorm"
)

type UserDb interface {
	CreateUser(models.User) (uint64, error)
	GetUserById(uint64) (models.User, error)
	GetUsers() ([]models.User, error)
	UserLogin(resources.UserCredential) (models.User, error)
	DeleteUserById(uint64) error
}

type userDb struct {
	db *gorm.DB
}

func NewUserDb() *userDb {
	return &userDb{
		db: config.Conn,
	}
}

func (d *userDb) CreateUser(user models.User) (uint64, error) {
	funcdesc := "CreateUser"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Save(&user)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return user.ID, err
	}
	log.Println("exit " + funcdesc)
	return user.ID, nil
}

func (d *userDb) GetUserById(userId uint64) (user models.User, err error) {
	funcdesc := "GetUserById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", userId).Find(&user)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return user, err
	} else if user.ID == 0 {
		return user, errors.New("user id doesn't exist")
	}
	log.Println("exit " + funcdesc)
	return user, nil
}

func (d *userDb) GetUsers() (users []models.User, err error) {
	funcdesc := "GetUsers"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Find(&users)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return users, err
	}
	log.Println("exit " + funcdesc)
	return users, nil
}

func (d *userDb) UserLogin(user resources.UserCredential) (output models.User, err error) {
	funcdesc := "UserLogin"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("email=?", user.UserName).Find(&output)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return output, err
	}
	fmt.Println(output)

	log.Println("exit " + funcdesc)
	return output, nil
}

func (d *userDb) DeleteUserById(userId uint64) error {
	funcdesc := "DeleteUserById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", userId).Delete(models.User{})
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
