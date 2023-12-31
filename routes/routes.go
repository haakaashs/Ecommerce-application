package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/config"
)

func Start() {
	buildHandler()
	router := gin.Default()
	router.Use(gin.Logger())

	// Use the Recovery middleware to handle panics
	router.Use(gin.Recovery())

	// user APIs
	user := router.Group("/user")
	{
		// for simulation password is stored directly
		// create and update user
		user.POST("/save", userHandler.CreateUser)

		// get user by id
		user.GET("/:user_id", userHandler.GetUserById)

		// get all users
		user.GET("/get-all", userHandler.GetUsers)

		// user login
		user.POST("/login", userHandler.UserLogin)

		// delete user by id
		user.DELETE("/:user_id", userHandler.DeleteUserById)
	}

	// product APIs
	product := router.Group("/product")
	{
		// create and update product
		product.POST("/save", productHandler.CreateProduct)

		// get product by id
		product.GET("/:product_id", productHandler.GetProductById)

		// get all the products
		product.GET("/get-all", productHandler.GetProducts)

		// delete product by id
		product.DELETE("/:product_id", productHandler.DeleteProductById)
	}

	// cart APIs
	cart := router.Group("/cart")
	{
		// create and update cart
		cart.POST("/save", cartHandler.CreateCart)

		// get cart by ID
		cart.GET("/:user_id", cartHandler.GetCartById)

		// delete cart by ID
		cart.DELETE("/:user_id", cartHandler.DeleteCartById)
	}

	// order APIs
	order := router.Group("/order")
	{
		// create order
		order.POST("/save", orderHandler.CreateOrder)

		// get order by ID
		order.GET("/:order_id", orderHandler.GetOrderById)

		// update the order status to cancelled
		order.PUT("/status/:order_id", orderHandler.UpdateOrderStatus)
	}

	// Lestening on port 8081
	log.Println("Starting server ...............")
	router.Run(config.Config.ServerPort)

}
