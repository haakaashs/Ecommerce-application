package routes

import (
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/handler"
	"github.com/haakaashs/antino-labs/service"
)

var (
	userDb    database.UserDb
	productDb database.ProductDb
	orderDb   database.OrderDb
	cartDb    database.CartDB
)

var (
	userService    service.UserService
	productService service.ProductService
	orderService   service.OrderService
	cartService    service.CartService
)

var (
	userHandler    handler.UserHandler
	productHandler handler.ProductHandler
	orderHandler   handler.OrderHandler
	cartHandler    handler.CartHandler
)

func buildDB() {
	userDb = database.NewUserDb()
	productDb = database.NewProductDb()
	orderDb = database.NewOrderDb()
	cartDb = database.NewCartDB()
}

func buildService() {
	userService = service.NewUserService(userDb)
	productService = service.NewProductService(productDb)
	orderService = service.NewOrderService(orderDb)
	cartService = service.NewCartService(cartDb, productDb)
}

func buildHandler() {
	buildDB()
	buildService()
	userHandler = handler.NewUserHandler(userService)
	productHandler = handler.NewProductHandler(productService)
	orderHandler = handler.NewOrderHandler(orderService)
	cartHandler = handler.NewCartHandler(cartService)
}
