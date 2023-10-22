package service

import (
	"log"

	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
)

type CartService interface {
	CreateCart(cart models.Cart) (uint64, error)
	GetCartById(cartId uint64) (cart models.Cart, err error)
	DeleteCartById(cartId uint64) error
}

type cartService struct {
	cartDB database.CartDB
}

func NewCartService(cartDB database.CartDB) *cartService {
	return &cartService{
		cartDB: cartDB,
	}
}

func (s *cartService) CreateCart(cart models.Cart) (uint64, error) {
	funcdesc := "CreateCart"
	log.Println("enter service" + funcdesc)

	cartId, err := s.cartDB.CreateCart(cart)
	if err != nil {
		return cartId, err
	}

	log.Println("exit " + funcdesc)
	return cartId, nil
}

func (s *cartService) GetCartById(cartId uint64) (cart models.Cart, err error) {
	funcdesc := "GetCartById"
	log.Println("enter service" + funcdesc)

	cart, err = s.cartDB.GetCartById(cartId)
	if err != nil {
		return cart, err
	}

	log.Println("exit " + funcdesc)
	return cart, nil
}

func (s *cartService) DeleteCartById(cartId uint64) error {
	funcdesc := "DeleteCartById"
	log.Println("enter service" + funcdesc)

	err := s.cartDB.DeleteCartById(cartId)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
