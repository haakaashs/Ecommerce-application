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

	// user APIs
	user := router.Group("/user")
	{
		// create and update user
		user.POST("/create", userHandler.CreateUser)

		// get user by id
		user.GET("/{user_id}", userHandler.GetUserById)

		// get all users
		user.GET("/get-all", userHandler.GetUsers)

		// user login
		user.POST("/login", userHandler.UserLogin)

		// delete user by id
		user.DELETE("/{user_id}", userHandler.DeleteUserById)
	}

	// product APIs
	product := router.Group("/product")
	{
		// create and update product
		product.POST("/create", productHandler.CreateProduct)

		// get product by id
		product.GET("/{product_id}", productHandler.GetProductById)

		// get all the products
		product.GET("/get-all", productHandler.GetProducts)

		// delete product by id
		product.POST("/{product_id}", productHandler.DeleteProductById)
	}

	// cart APIs
	cart := router.Group("/cart")
	{
		// create and update cart
		cart.POST("/create", cartHandler.CreateCart)

		// get cart by ID
		cart.GET("/{card_id}", cartHandler.GetCartById)

		// delete cart by ID
		cart.DELETE("/{cart_id}", cartHandler.DeleteCartById)
	}

	// order APIs
	order := router.Group("/order")
	{
		// create order
		order.POST("/create")

		// get order by ID
		order.GET("/{order_id}")

		// update the order status
		order.PUT("/update")
	}

	// Lestening on port 8081
	log.Println("Starting server ...............")
	router.Run(config.Config.ServerPort)

}
