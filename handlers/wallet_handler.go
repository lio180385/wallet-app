package handlers

import (
	"net/http"
	"wallet-app/database"
	"wallet-app/models"
	"github.com/gin-gonic/gin"
)

type WithdrawRequest struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}

// Withdraw API
func Withdraw(c *gin.Context) {
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if user.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	user.Balance -= req.Amount
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "withdrawal successful",
		"balance": user.Balance,
	})
}

// Balance Inquiry API
func GetBalance(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
		"name":    user.Name,
		"balance": user.Balance,
	})
}
