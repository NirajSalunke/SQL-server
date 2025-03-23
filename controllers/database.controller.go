package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/sql-maker/config"
	"www.github.com/NirajSalunke/sql-maker/models"
)

func CreateDatabase(c *gin.Context) {
	var newDB models.Database

	if err := c.ShouldBindJSON(&newDB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input data!",
			"error":   err.Error(),
		})
		return
	}

	var currUser models.User
	if res := config.DB.First(&currUser, newDB.UserID); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
			"error":   res.Error.Error(),
		})
		return
	}

	newDB.UserID = currUser.ID
	if err := config.DB.Create(&newDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create database!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"message":  "Database created successfully!",
		"database": newDB,
	})
}
