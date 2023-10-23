package service

import (
	"errors"
	"log"

	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/utils"
)

type CartService interface {
	CreateCart(cart resources.CartResource) (uint64, error)
	GetCartById(cartId uint64) (cart models.Cart, err error)
	DeleteCartById(cartId uint64) error
}

type cartService struct {
	cartDB    database.CartDB
	productDb database.ProductDb
}

func NewCartService(cartDB database.CartDB, productDb database.ProductDb) *cartService {
	return &cartService{
		cartDB:    cartDB,
		productDb: productDb,
	}
}

func (s *cartService) CreateCart(cart resources.CartResource) (uint64, error) {
	funcdesc := "CreateCart"
	log.Println("enter service" + funcdesc)

	// calculate the total cart amount for the update call
	err := s.calculateCartValue(&cart)
	if err != nil {
		return cart.ID, err
	}

	// resource to model conversion
	cartM, err := utils.ResourceToModel(&cart)
	if err != nil {
		return cart.ID, err
	}

	cartId, err := s.cartDB.CreateCart(cartM)
	if err != nil {
		return cartId, err
	}

	log.Println("exit " + funcdesc)
	return cartId, nil
}

func (s *cartService) calculateCartValue(cart *resources.CartResource) error {

	var (
		TotalAmount float64
		count       uint
	)

	for _, product := range cart.CartProducts {

		//  maximun order quantity check
		if product.ProductQty > 10 {
			return errors.New("maximum product quantity should not exceed 10 for " + product.ProductName)
		}

		// product uint price check
		productDetails, err := s.productDb.GetProductById(product.Id)
		if err != nil {
			return err
		} else if productDetails.Price != product.ProductUnitPrice {
			return errors.New("incorrect product uint price for " + product.ProductName)
		}

		// product total price check
		ProductTotalPrice := float64(product.ProductQty) * productDetails.Price
		if product.ProductTotalPrice != ProductTotalPrice {
			return errors.New("incorrect product total price for " + product.ProductName)
		}
		TotalAmount += ProductTotalPrice

		// 10% discount if 3 premium products are added
		if productDetails.ProductCategory == constants.PREMIMUM {
			count += 1
		}
	}

	if count > 3 {
		discount := TotalAmount * 0.1
		cart.TotalAmount = TotalAmount - discount
		return nil
	}

	cart.TotalAmount = TotalAmount
	return nil
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
