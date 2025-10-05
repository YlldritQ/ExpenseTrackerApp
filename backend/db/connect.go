package db

import (
	"log"
	"os"

	"github.com/YlldritQ/ExpenseTrackerApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	var DB *gorm.DB

func Connect() {
	
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Expense{})

	DB = db

	log.Println("Database connection established")
}