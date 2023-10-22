package handler

import "github.com/haakaashs/antino-labs/service"

type CartHandler interface {
}

type cartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *cartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}
