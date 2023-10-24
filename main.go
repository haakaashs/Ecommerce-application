package main

import (
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/routes"
)

func main() {

	// initialize validator
	config.InitializeValidator()

	// initialize database
	config.InitializeDB("./config.env")

	// start application
	routes.Start()
}
