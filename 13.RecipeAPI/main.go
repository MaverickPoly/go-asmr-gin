package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ConnectDB()
	app := gin.Default()
	RegisterRoutes(app)
	app.Run(":8000")
}
