package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/sql-maker/config"
	"www.github.com/NirajSalunke/sql-maker/models"
	"www.github.com/NirajSalunke/sql-maker/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectToDatabase()
	config.SetupGeminiClient()
	models.MigrateModels()
}
func main() {
	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run(os.Getenv("PORT"))
}
