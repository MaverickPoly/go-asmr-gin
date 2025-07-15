package main

import (
	"08.BookManagement/config"
	"08.BookManagement/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()
	app := gin.Default()
	routes.HandleRoutes(app)

	app.Run(":8000")
}
