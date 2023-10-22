package handler

import "github.com/haakaashs/antino-labs/service"

type ProductHandle interface {
}

type productHandle struct {
	productService service.ProductService
}

func NewProductHandle(productService service.ProductService) *productHandle {
	return &productHandle{
		productService: productService,
	}
}
