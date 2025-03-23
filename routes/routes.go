package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	userGroup := r.Group("/user")
	databaseGroup := r.Group("/database")
	queryGroup := r.Group("/query")
	LoadUserRoutes(userGroup)
	LoadDatabaseRoutes(databaseGroup)
	LoadQueryRoutes(queryGroup)

}
