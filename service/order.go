package service

import (
	"log"

	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
)

type OrderService interface {
	CreateOrder(models.Order) (uint64, error)
	GetOrderById(uint64) (models.Order, error)
	UpdateOrderStatus(resources.OrderStatusUpdate) error
}

type orderService struct {
	orderDB database.OrderDb
}

func NewOrderService(orderDB database.OrderDb) *orderService {
	return &orderService{
		orderDB: orderDB,
	}
}

func (s *orderService) CreateOrder(order models.Order) (uint64, error) {
	funcdesc := "CreateOrder"
	log.Println("enter service" + funcdesc)

	orderId, err := s.orderDB.CreateOrder(order)
	if err != nil {
		return orderId, err
	}

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
