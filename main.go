package main

import (
	"ecommerce-backend/database"
	"ecommerce-backend/router"
)

func main() {

	// Read config from file

	// Init Database
	database.Init()
	// Init Router
	r := router.Init()

	// Start Router
	r.Start()
}
