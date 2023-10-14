package main

import (
	"task-5-pbi-btpns-muhammad-ikhwan-fathulloh/database"
	"task-5-pbi-btpns-muhammad-ikhwan-fathulloh/router"
)

func main() {
	// Connect to database
	database.Connect()
	// Run migrations
	database.Migrate()

	// Setup the router
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}
