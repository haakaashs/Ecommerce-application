package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/utils"
)

type CartService interface {
	CreateCart(resources.CartResource) (uint64, error)
	GetCartById(uint64) (resources.CartResource, error)
	DeleteCartById(uint64) error
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

	cart.Qty = uint(len(cart.CartProducts))

	// resource to model conversion
	cartM := utils.ResourceToModel(cart)

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

	for index, product := range cart.CartProducts {

		// validate product details
		err := config.Validate.Struct(product)
		if err != nil {
			return errors.New("error: " + err.Error())
		}

		//  maximun order quantity check
		if product.ProductQty > 10 {
			return errors.New("maximum product quantity should not exceed 10 for " + product.ProductName)
		}

		// get product details
		productDetails, err := s.productDb.GetProductById(product.ProductId)
		if err != nil {
			return err
		}

		// product total price and cark total price calculated
		ProductTotalPrice := float64(product.ProductQty) * productDetails.Price
		cart.CartProducts[index].ProductUnitPrice = productDetails.Price
		cart.CartProducts[index].ProductTotalPrice = ProductTotalPrice
		TotalAmount += ProductTotalPrice

		// 10% discount if 3 premium products are added
		if productDetails.ProductCategory == constants.PREMIMUM {
			count += 1
		}
	}
	fmt.Println("total: ", count)

	if count >= 3 {
		cart.TotalAmount = TotalAmount * 0.9
		cart.Discount = TotalAmount * 0.1
		fmt.Println("jill: ", cart.Discount, ":", TotalAmount*0.1)
		return nil
	}

	cart.TotalAmount = TotalAmount
	return nil
}

func (s *cartService) GetCartById(userId uint64) (cart resources.CartResource, err error) {
	funcdesc := "GetCartById"
	log.Println("enter service" + funcdesc)

	cartM, err := s.cartDB.GetCartById(userId)
	if err != nil {
		return cart, err
	}

	cart, err = utils.ModelToResource(cartM)
	if err != nil {
		return cart, err
	}

	log.Println("exit " + funcdesc)
	return cart, nil
}

func (s *cartService) DeleteCartById(userId uint64) error {
	funcdesc := "DeleteCartById"
	log.Println("enter service" + funcdesc)

	err := s.cartDB.DeleteCartById(userId)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
