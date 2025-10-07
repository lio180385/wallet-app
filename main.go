package main

import (
	"wallet-app/database"
	"wallet-app/handlers"
	"wallet-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect DB
	database.ConnectDatabase()

	// Seed data user (optional)
	database.DB.FirstOrCreate(&models.User{ID: 1}, models.User{Name: "Leo", Balance: 100000})

	// Routes
	r.POST("/withdraw", handlers.Withdraw)
	r.GET("/balance/:id", handlers.GetBalance)

	// Run server
	r.Run(":8080")
}
