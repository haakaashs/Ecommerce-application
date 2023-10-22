package service

import "github.com/haakaashs/antino-labs/database"

type CartService interface {
}

type cartService struct {
	cartDB database.CartDB
}

func NewCartService(cartDB database.CartDB) *cartService {
	return &cartService{
		cartDB: cartDB,
	}
}
