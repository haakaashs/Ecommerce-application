package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/service"
)

type CartHandler interface {
	CreateCart(ctx *gin.Context)
	GetCartById(ctx *gin.Context)
	DeleteCartById(ctx *gin.Context)
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

	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	cartId, err := strconv.Atoi(ctx.Param("cart_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := h.cartService.GetCartById(uint64(cartId))
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

	cartId, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.cartService.DeleteCartById(uint64(cartId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Successfully Deleted"})
}
