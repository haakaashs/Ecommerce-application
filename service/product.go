package service

import (
	"log"

	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
)

type ProductService interface {
	CreateProduct(models.Product) error
	GetProductById(uint64) (models.Product, error)
	GetProducts() ([]models.Product, error)
	DeleteProductById(uint64) error
}

type productService struct {
	productDb database.ProductDb
}

func NewProductService(productDb database.ProductDb) *productService {
	return &productService{
		productDb: productDb,
	}
}

func (s *productService) CreateProduct(product models.Product) error {
	funcdesc := "CreateProduct"
	log.Println("enter service" + funcdesc)

	// while login check user permission for add,edit,delete product
	// assuming that user permission is obtained when login

	err := s.productDb.CreateProduct(product)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}

func (s *productService) GetProductById(productId uint64) (product models.Product, err error) {
	funcdesc := "GetProductById"
	log.Println("enter service" + funcdesc)

	product, err = s.productDb.GetProductById(productId)
	if err != nil {
		return product, err
	}

	log.Println("exit " + funcdesc)
	return product, nil
}

func (s *productService) GetProducts() (products []models.Product, err error) {
	funcdesc := "GetProducts"
	log.Println("enter service" + funcdesc)

	products, err = s.productDb.Getproducts()
	if err != nil {
		return products, err
	}

	log.Println("exit " + funcdesc)
	return products, nil
}

func (s *productService) DeleteProductById(productId uint64) error {
	funcdesc := "DeleteProductById"
	log.Println("enter service" + funcdesc)

	err := s.productDb.DeleteProductById(productId)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
