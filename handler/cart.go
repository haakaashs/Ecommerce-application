package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/service"
)

type CartHandler interface {
	CreateCart(*gin.Context)
	GetCartById(*gin.Context)
	DeleteCartById(*gin.Context)
}

type cartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *cartHandler {
	return &cartHandler{
		cartService: cartService,
	}
}

func (h *cartHandler) CreateCart(ctx *gin.Context) {
	funcdesc := "CreateCart"
	log.Println("enter handeler" + funcdesc)

	var cart resources.CartResource
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.Validate.Struct(cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cartId, err := h.cartService.CreateCart(cart)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"cart_id": cartId})
}

func (h *cartHandler) GetCartById(ctx *gin.Context) {
	funcdesc := "GetCartById"
	log.Println("enter handeler" + funcdesc)

	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.cartService.GetCartById(uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, cart)
}

func (h *cartHandler) DeleteCartById(ctx *gin.Context) {
	funcdesc := "DeleteCartById"
	log.Println("enter handeler" + funcdesc)

	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.cartService.DeleteCartById(uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Successfully Deleted"})
}
