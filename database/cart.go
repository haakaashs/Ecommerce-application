package database

import (
	"errors"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"gorm.io/gorm"
)

type CartDB interface {
	CreateCart(models.Cart) (uint64, error)
	GetCartById(uint64) (models.Cart, error)
	DeleteCartById(uint64) error
	GetCartByUserId(UserId uint64) (cart models.Cart, err error)
}

type cartDB struct {
	db *gorm.DB
}

func NewCartDB() *cartDB {
	return &cartDB{
		db: config.Conn,
	}
}

func (d *cartDB) CreateCart(cart models.Cart) (uint64, error) {
	funcdesc := "CreateCart"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Create(&cart)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return cart.ID, err
	}

	log.Println("exit " + funcdesc)
	return cart.ID, nil
}

func (d *cartDB) GetCartById(userId uint64) (cart models.Cart, err error) {
	funcdesc := "GetCartById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("user_id=?", userId).Find(&cart)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return cart, err
	} else if cart.ID == 0 {
		return cart, errors.New("user id doesn't exist")
	}
	log.Println("exit " + funcdesc)
	return cart, nil
}

func (d *cartDB) DeleteCartById(userId uint64) error {
	funcdesc := "DeleteCartById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("user_id=?", userId).Delete(models.Cart{})
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}

func (d *cartDB) GetCartByUserId(UserId uint64) (cart models.Cart, err error) {
	funcdesc := "GetCartByUserId"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("user_id=?", UserId).Find(&cart)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return cart, err
	}
	log.Println("exit " + funcdesc)
	return cart, nil
}
