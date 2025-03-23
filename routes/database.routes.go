package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/sql-maker/controllers"
)

func LoadDatabaseRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CreateDatabase)
}
