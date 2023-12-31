package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/service"
)

type OrderHandler interface {
	CreateOrder(*gin.Context)
	GetOrderById(*gin.Context)
	UpdateOrderStatus(*gin.Context)
}

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (h *orderHandler) CreateOrder(ctx *gin.Context) {
	funcdesc := "CreateOrder"
	log.Println("enter handeler" + funcdesc)

	var order resources.OrderResource
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.Validate.Struct(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderId, err := h.orderService.CreateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"order_id": orderId})
}

func (h *orderHandler) GetOrderById(ctx *gin.Context) {
	funcdesc := "GetOrderById"
	log.Println("enter handeler" + funcdesc)

	orderId, err := strconv.Atoi(ctx.Param("order_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.GetOrderById(uint64(orderId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, order)
}

func (h *orderHandler) UpdateOrderStatus(ctx *gin.Context) {
	funcdesc := "UpdateOrderStatus"
	log.Println("enter handeler" + funcdesc)

	orderId, err := strconv.Atoi(ctx.Param("order_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := resources.OrderStatusUpdate{
		OrderId:     uint64(orderId),
		OrderStatus: constants.CANCELLED,
	}

	err = h.orderService.UpdateOrderStatus(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Successfully Updated"})
}
