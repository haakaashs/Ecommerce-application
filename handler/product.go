package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/service"
)

type ProductHandler interface {
	CreateProduct(*gin.Context)
	GetProductById(*gin.Context)
	GetProducts(*gin.Context)
	DeleteProductById(*gin.Context)
}

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) CreateProduct(ctx *gin.Context) {
	funcdesc := "CreateProduct"
	log.Println("enter handeler" + funcdesc)

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.Validate.Struct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productId, err := h.productService.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"product_id": productId})
}

func (h *productHandler) GetProductById(ctx *gin.Context) {
	funcdesc := "GetProductById"
	log.Println("enter handeler" + funcdesc)

	productId, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.productService.GetProductById(uint64(productId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) GetProducts(ctx *gin.Context) {
	funcdesc := "GetProducts"
	log.Println("enter handeler" + funcdesc)

	products, err := h.productService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, products)
}

func (h *productHandler) DeleteProductById(ctx *gin.Context) {
	funcdesc := "DeleteProductById"
	log.Println("enter handeler" + funcdesc)

	productId, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.productService.DeleteProductById(uint64(productId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Deletion Successful"})
}
