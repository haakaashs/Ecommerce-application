package service

import (
	"log"
	"time"

	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/utils"
)

type OrderService interface {
	CreateOrder(resources.OrderResource) (uint64, error)
	GetOrderById(uint64) (resources.OrderResource, error)
	UpdateOrderStatus(resources.OrderStatusUpdate) error
}

type orderService struct {
	orderDB   database.OrderDb
	cartDb    database.CartDB
	productDb database.ProductDb
}

func NewOrderService(orderDB database.OrderDb, cartDb database.CartDB, productDb database.ProductDb) *orderService {
	return &orderService{
		orderDB:   orderDB,
		cartDb:    cartDb,
		productDb: productDb,
	}
}

func (s *orderService) CreateOrder(order resources.OrderResource) (orderId uint64, err error) {
	funcdesc := "CreateOrder"
	log.Println("enter service" + funcdesc)

	cart, err := s.cartDb.GetCartByUserId(order.UserID)
	if err != nil {
		return 0, err
	}

	orderM := models.Order{
		UserID:        order.UserID,
		OrderStatus:   constants.PLACED,
		OrderProducts: cart.CartProducts,
		Discount:      cart.Discount,
		OrderValue:    cart.TotalAmount,
		IsActive:      &[]bool{true}[0],
	}

	orderId, err = s.orderDB.CreateOrder(orderM)
	if err != nil {
		return orderId, err
	}

	// inventory update
	go func(cart models.Cart) {
		cartResource, _ := utils.CartModelToResource(cart)
		for index := range cartResource.CartProducts {
			product, _ := s.productDb.GetProductById(cartResource.CartProducts[index].ProductId)
			product.InventoryQty = product.InventoryQty - cartResource.CartProducts[index].ProductQty
			s.productDb.CreateProduct(product)
		}
	}(cart)

	// order status
	go func(ordId uint64) {
		time.Sleep(time.Second * 5)
		for _, status := range []string{constants.DISPATCHED, constants.COMPLETED} {
			time.Sleep(time.Second * 5)
			s.orderDB.UpdateOrderStatus(resources.OrderStatusUpdate{OrderId: ordId, OrderStatus: status})
		}
	}(orderId)

	log.Println("exit " + funcdesc)
	return orderId, nil
}

func (s *orderService) GetOrderById(orderId uint64) (orderR resources.OrderResource, err error) {
	funcdesc := "GetOrderById"
	log.Println("enter service" + funcdesc)

	order, err := s.orderDB.GetOrderById(orderId)
	if err != nil {
		return orderR, err
	}

	orderR, err = utils.OrderModelToResource(order)
	if err != nil {
		return orderR, err

	}

	log.Println("exit " + funcdesc)
	return orderR, nil
}

func (s *orderService) UpdateOrderStatus(orderUpdate resources.OrderStatusUpdate) error {
	funcdesc := "UpdateOrderStatus"
	log.Println("enter service" + funcdesc)

	err := s.orderDB.UpdateOrderStatus(orderUpdate)
	if err != nil {
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
