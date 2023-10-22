package service

import "github.com/haakaashs/antino-labs/database"

type ProductService interface {
}

type productService struct {
	productDB database.ProductDb
}

func NewProductService(productDB database.ProductDb) *productService {
	return &productService{
		productDB: productDB,
	}
}
