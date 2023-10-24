package database

import (
	"errors"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"gorm.io/gorm"
)

type ProductDb interface {
	CreateProduct(models.Product) (uint64, error)
	GetProductById(uint64) (models.Product, error)
	Getproducts() ([]models.Product, error)
	DeleteProductById(uint64) error
}

type productDb struct {
	db *gorm.DB
}

func NewProductDb() *productDb {
	return &productDb{
		db: config.Conn,
	}
}

func (d *productDb) CreateProduct(product models.Product) (uint64, error) {
	funcdesc := "CreateProduct"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Save(&product)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return product.ID, err
	}
	log.Println("exit " + funcdesc)
	return product.ID, nil
}

func (d *productDb) GetProductById(productId uint64) (product models.Product, err error) {
	funcdesc := "GetProductById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", productId).Find(&product)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return product, err
	} else if product.ID == 0 {
		return product, errors.New("product id doesn't exist")
	}
	log.Println("exit " + funcdesc)
	return product, nil
}

func (d *productDb) Getproducts() (products []models.Product, err error) {
	funcdesc := "Getproducts"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Find(&products)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return products, err
	}
	log.Println("exit " + funcdesc)
	return products, nil
}

func (d *productDb) DeleteProductById(productId uint64) error {
	funcdesc := "DeleteProductById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", productId).Delete(models.Product{})
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}

// func (d *productDb) GetProductsByIds(productIds []uint64) (map[uint64]float64, error) {
// 	funcdesc := "GetProductsByIds"
// 	log.Println("enter DB" + funcdesc)

// 	productMap := map[uint64]float64{}
// 	result := d.db.Debug().Where("id in", productIds).Find(&productMap)
// 	if err := result.Error; err != nil {
// 		log.Println("error in DB query: ", err.Error())
// 		return productMap, err
// 	}

// 	log.Println("exit " + funcdesc)
// 	return productMap, nil
// }
