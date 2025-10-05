package main

import (
	"log"
	"os"

	"github.com/YlldritQ/ExpenseTrackerApp/config"
	"github.com/YlldritQ/ExpenseTrackerApp/db"
	"github.com/YlldritQ/ExpenseTrackerApp/routes"
)

func main() {
	config.LoadEnv()
	db.Connect()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting server on port %s 🚀", port)
	r := routes.SetupRouter()
	r.Run(":" + port)
}