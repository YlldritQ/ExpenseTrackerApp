package handlers

import (
	"net/http"
	"time"

	"github.com/YlldritQ/ExpenseTrackerApp/db"
	"github.com/YlldritQ/ExpenseTrackerApp/models"
	"github.com/gin-gonic/gin"
)

func GetExpenses(c *gin.Context) {
	var expenses []models.Expense
	if err := db.DB.Preload("User").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func GetUserExpenses(c *gin.Context) {
	var expenses []models.Expense
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	
	uid := userId.(uint)
	if err := db.DB.Where("user_id = ?", uid).Preload("User").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func AddExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	uid := userId.(uint)
	expense.UserID = &uid
	expense.Date = time.Now()
	if err := db.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Expense{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted"})
}