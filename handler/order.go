package handler

import "github.com/haakaashs/antino-labs/service"

type OrderHandler interface {
}

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}
