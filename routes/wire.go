package routes

import (
	"github.com/haakaashs/antino-labs/database"
	"github.com/haakaashs/antino-labs/handler"
	"github.com/haakaashs/antino-labs/service"
)

var (
	basicDB   database.BasicDB
	userDb    database.UserDb
	productDb database.ProductDb
	orderDb   database.OrderDb
	cartDb    database.CartDB
)

var (
	basicService   service.BasicService
	userService    service.UserService
	productService service.ProductService
	orderService   service.OrderService
	cartService    service.CartService
)

var (
	basicHandler   handler.BasicHandler
	userHandler    handler.UserHandle
	productHandler handler.ProductHandle
	orderHandler   handler.OrderHandler
	cartHandler    handler.CartHandler
)

func buildDB() {
	basicDB = database.NewBasicDB()
	userDb = database.NewUserDb()
	productDb = database.NewProductDb()
	orderDb = database.NewOrderDb()
	cartDb = database.NewCartDB()
}

func buildService() {
	basicService = service.NewBasicService(basicDB)
	userService = service.NewUserService(userDb)
	productService = service.NewProductService(productDb)
	orderService = service.NewOrderService(orderDb)
	cartService = service.NewCartService(cartDb)
}

func buildHandler() {
	buildDB()
	buildService()
	basicHandler = handler.NewBasicHandler(basicService)
	userHandler = handler.NewUserHandle(userService)
	productHandler = handler.NewProductHandle(productService)
	orderHandler = handler.NewOrderHandler(orderService)
	cartHandler = handler.NewCartHandler(cartService)
}
