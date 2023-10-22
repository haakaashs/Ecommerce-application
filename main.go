package main

import (
	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/routes"
)

func main() {
	config.InitializeDB("./config.env")
	routes.Start()
}
