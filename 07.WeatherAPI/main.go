package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := gin.Default()
	SetupRoutes(app)

	app.Run(":8080")
}
