package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/sql-maker/client"
	"www.github.com/NirajSalunke/sql-maker/config"
	"www.github.com/NirajSalunke/sql-maker/models"
)

func GetQuery(c *gin.Context) {
	var currQuery *models.QueryRequest
	if err := c.ShouldBindJSON(&currQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input data!",
			"error":   err.Error(),
		})
		return
	}
	var queryString string
	var err error
	if queryString, err = client.NaturalTextToSQL(currQuery); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to execute query!",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println("Query:- ", queryString)
	conversation := models.Conversation{
		UserInput:  currQuery.NaturalText,
		AiOutput:   queryString,
		DatabaseID: currQuery.DatabaseID,
	}

	if err := config.DB.Create(&conversation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to save conversation!",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Query Generated successfully!",
		"query":   queryString,
	})
}
