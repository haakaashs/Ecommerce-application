package service

import (
	"log"
	"time"
	// "time"

	// "github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/utils"
	// "github.com/haakaashs/antino-labs/utils"
)

type OrderService interface {
	CreateOrder(models.Order) (uint64, error)
	GetOrderById(uint64) (models.Order, error)
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

func (s *orderService) CreateOrder(order models.Order) (orderId uint64, err error) {
	funcdesc := "CreateOrder"
	log.Println("enter service" + funcdesc)

	cart, err := s.cartDb.GetCartByUserId(order.UserID)
	if err != nil {
		return 0, err
	}

	order.OrderProducts = cart.CartProducts
	order.OrderStatus = constants.PLACED
	order.IsActive = true
	orderId, err = s.orderDB.CreateOrder(order)
	if err != nil {
		return orderId, err
	}

	// inventory update
	go func(cart models.Cart) {
		cartResource, _ := utils.ModelToResource(cart)
		for _, val := range cartResource.CartProducts {
			product, _ := s.productDb.GetProductById(val.Id)
			product.InventoryQty = product.InventoryQty - val.ProductQty
			s.productDb.CreateProduct(product)
		}
	}(cart)

	// order status
	go func(ordId uint64) {
		time.Sleep(time.Second * 5)
		for _, status := range []string{constants.DISPATCHED, constants.COMPLETED} {
			time.Sleep(time.Second * 5)
			s.UpdateOrderStatus(resources.OrderStatusUpdate{OrderId: ordId, OrderStatus: status})

		}
	}(order.ID)

	log.Println("exit " + funcdesc)
	return orderId, nil
}

func (s *orderService) GetOrderById(orderId uint64) (models.Order, error) {
	funcdesc := "GetOrderById"
	log.Println("enter service" + funcdesc)

	order, err := s.orderDB.GetOrderById(orderId)
	if err != nil {
		return order, err
	}

	log.Println("exit " + funcdesc)
	return order, nil
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
