package routes

import (
	"github.com/YlldritQ/ExpenseTrackerApp/handlers"
	"github.com/YlldritQ/ExpenseTrackerApp/middleware"
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

	AuthRoutes(r)
	// Protected routes

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("expenses", handlers.GetExpenses)
		protected.GET("expenses/user", handlers.GetUserExpenses)
		protected.POST("expenses", handlers.AddExpense)
		protected.DELETE("expenses/:id", handlers.DeleteExpense)
	}
	
	return r
}