package service

import "github.com/haakaashs/antino-labs/database"

type OrderService interface {
}

type orderService struct {
	orderDB database.OrderDb
}

func NewOrderService(orderDB database.OrderDb) *orderService {
	return &orderService{
		orderDB: orderDB,
	}
}
