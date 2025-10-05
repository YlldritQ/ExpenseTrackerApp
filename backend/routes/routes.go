package routes

import (
	"github.com/YlldritQ/ExpenseTrackerApp/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*", "http://localhost:5173"}, // or "*"
    AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
}))

	r.GET("expenses", handlers.GetExpenses)
	r.POST("expenses", handlers.AddExpense)
	r.DELETE("expenses/:id", handlers.DeleteExpense)

	return r
}